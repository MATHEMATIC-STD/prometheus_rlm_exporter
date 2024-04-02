package statics

const REG_TOTAL = `(?m)((?P<product>\w+)\s(?P<version>v\d+\.\d+).+\n\s+count:\s(?P<count>\d+))`
const REG_USAGE = `(?m)((?P<product>\w+)\s(?P<version>v\d+\.\d+):\W(?P<user>[\w\-\.]+)\W(?P<worker>[\w\-\d]+))`
