#!/usr/bin/env python
import smtplib

server = smtplib.SMTP(host='localhost', port=2525)
server.set_debuglevel(1)
server.sendmail("debug@example.com", "factory-minasan@example.com", """From: robert
Subject: super mail

This is fine""")
server.quit()