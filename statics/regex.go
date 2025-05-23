package statics

const REG_TOTAL = `(?m)((?P<product>\w+)\s(?P<version>v[\d\.]+).+\n\s+count:\s(?P<count>\d+))`
const REG_USAGE = `(?m)((?P<product>\w+)\s(?P<version>v[\d\.]+):\W(?P<user>[\w\-\.]+)\W(?P<computer>[\w\-\d]+))`
