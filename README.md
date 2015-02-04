naisho [![Build Status](https://travis-ci.org/moznion/naisho.svg)](https://travis-ci.org/moznion/naisho)
==

Send short encrypted message by GitHub ID

Description
--

Encrypt the short message by a ssh-rsa public key on GitHub,
and send it by a mail to user's email address which is written on GitHub profile page.

It requires gmail account and authentication information, because
this program sends email via gmail.

Usage
--

Put the configuration file of gmail on your home directory;

```
cat <<EOS >$HOME/.naisho
address: example@gmail.com
password: PASSWORD
EOS
```

And run the command;

```
$ naisho <Target GitHub ID> <Message>
```

Then it sends the email to target GitHub user's mail address (it is on profile page)
with an attachment which contains encrypted message.

When you want to read a received attachment, you just execute following command;

```
$ openssl rsautl -decrypt -inkey <Your Secret Key> -in <Attachment>
```

Note
--

This program gets ssh-rsa public keys from [https://github.com/user-id.keys](https://github.com/user-id.keys)
and gets a target email address from [https://github.com/user-id](https://github.com/user-id)

Length of the secret message differs depending on strength of the used key.

Author
--

moznion (<moznion@gmail.com>)

License
--

MIT

