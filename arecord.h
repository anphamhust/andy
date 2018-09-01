#ifndef _ARECORD_H_
#define  _ARECORD_H_

// Record to a file
// Audio file is a 16-bit signed little-endian encoded
// with a sample rate of 16000.
int record_to_file(char * filename, int duration);

// Record to a buffer
int record_to_buffer(unsigned char * buffer, int duration);

#endif
