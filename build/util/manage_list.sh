#!/usr/bin/env bash
<<<<<<< HEAD
guodun1="$(../chain33-cli config config_tx -k token-finisher -o del -v 13KDST7ndBmzunGYAkCRabooWGRc95hM3W --paraName "user.p.starsunny.")"
guodun2="$(../chain33-cli config config_tx -k token-finisher -o add -v 1AtQcjpDYaJSxvAE3aacfJUknsFNb7eWLd --paraName "user.p.starsunny.")"
=======
guodun1=$(../chain33-cli config config_tx -k token-blacklist -o add -v BTY --paraName "user.p.$2.")
guodun2=$(../chain33-cli config config_tx -k token-finisher -o add -v "$1" --paraName "user.p.$2.")
>>>>>>> upstream/master
echo "cli wallet sign -d $guodun1 -a 1JmFaA6unrCFYEWPGRi7uuXY1KthTJxJEP -e 1h > d:/b.txt"
echo "cli wallet sign -d $guodun2 -a 1JmFaA6unrCFYEWPGRi7uuXY1KthTJxJEP -e 1h >> d:/b.txt"
