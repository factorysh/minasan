皆さん
======

Minasan is a SMTP server, linked to a Gitlab instance.

When you send a mail to user `{group}.{project}`, every gitlab's users of this project, with higher level than observator receive the mail.

Big picture
-----------

            +---------+    +---------+
    mail -> | Minasan | -> | Postfix +--+-> Alice
            +---+-----+    +---------+  |
                | REST                  +-> Bob
                v                       |
            +---------+                 +-> Charly
            | Gitlab  |
            +---------+
