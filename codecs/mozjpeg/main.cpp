#include <cstdint>
#include <cstdio>
#include <cstdlib>
#include <cstring>

#include "mozjpeg-4.1.1/jpeglib.h"
#include <emscripten.h>

extern "C" {
uint8_t *allocate(size_t size);
void deallocate(uint8_t *ptr);
uint64_t encode(uint8_t *img_in, int width, int height, int channels,
                uint8_t *img_out);
}

EMSCRIPTEN_KEEPALIVE
uint8_t *allocate(size_t size) { return new uint8_t[size]; }

EMSCRIPTEN_KEEPALIVE
void deallocate(uint8_t *ptr) { delete[] ptr; }

EMSCRIPTEN_KEEPALIVE
uint64_t encode(uint8_t *img_in, int width, int height, int channels,
                uint8_t *img_out) {
  struct jpeg_compress_struct cinfo;
  struct jpeg_error_mgr jerr;
  jpeg_create_compress(&cinfo);

  cinfo.err = jpeg_std_error(&jerr);
  cinfo.image_width = width;
  cinfo.image_height = height;
  cinfo.input_components = channels;
  cinfo.in_color_space = JCS_RGB;

  unsigned long buf_size = 0;
  uint8_t *buf;
  jpeg_set_defaults(&cinfo);
  jpeg_set_colorspace(&cinfo, JCS_YCbCr);
  jpeg_set_quality(&cinfo, 80, false);

  jpeg_mem_dest(&cinfo, &buf, &buf_size);
  jpeg_start_compress(&cinfo, true);

  JSAMPROW row_pointer[1];
  while (cinfo.next_scanline < cinfo.image_height) {
    row_pointer[0] = &img_in[cinfo.next_scanline * cinfo.image_width *
                             cinfo.input_components];
    jpeg_write_scanlines(&cinfo, row_pointer, 1);
  }

  jpeg_finish_compress(&cinfo);
  jpeg_destroy_compress(&cinfo);

  std::realloc(img_out, buf_size);
  std::memcpy(img_out, buf, buf_size);
  return buf_size;
}
