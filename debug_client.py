#!/usr/bin/env python
import smtplib

server = smtplib.SMTP(host='localhost', port=2525)
server.set_debuglevel(1)
server.sendmail("debug@example.com", "group-projet@example.com", "This is fine")
server.quit()