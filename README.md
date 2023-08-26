# Archiver
### What is it?
Archiver is a simple program for compressing or decompressing files.

All compressed files will save as .eva archive file.

## EVA file format

WIP

### Header

| Position (hex) | Size (bytes) |         Name         |                     Description                     |
|:---------------|:------------:|:--------------------:|:---------------------------------------------------:|
| 00             |      3       | File signature (EVA) | Contains 0x455641 (little-endian) as file signature |
| 03             |              |                      |                                                     |
| 0x             |              |                      |                                                     |
    