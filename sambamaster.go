/*


 */

// samba master, web UI for samba administrator
package main

// TODO: split to UI+AGENT

/*
// TODO: UI login secure by https://code.google.com/p/crypto-js/#HMAC https://code.google.com/archive/p/crypto-js/

HMAC
Keyed-hash message authentication codes (HMAC) is a mechanism for message authentication using cryptographic hash functions.

HMAC can be used in combination with any iterated cryptographic hash function.

<script src="http://crypto-js.googlecode.com/svn/tags/3.1.2/build/rollups/hmac-md5.js"></script>
<script src="http://crypto-js.googlecode.com/svn/tags/3.1.2/build/rollups/hmac-sha1.js"></script>
<script src="http://crypto-js.googlecode.com/svn/tags/3.1.2/build/rollups/hmac-sha256.js"></script>
<script src="http://crypto-js.googlecode.com/svn/tags/3.1.2/build/rollups/hmac-sha512.js"></script>
<script>
    var hash = CryptoJS.HmacMD5("Message", "Secret Passphrase");
    var hash = CryptoJS.HmacSHA1("Message", "Secret Passphrase");
    var hash = CryptoJS.HmacSHA256("Message", "Secret Passphrase");
    var hash = CryptoJS.HmacSHA512("Message", "Secret Passphrase");
</script>

Secret Passphrase is nonce get from server every time(login)
*/

/*
config:

// UI config
listen="0.0.0.0:8800"
autostart="true"
databasefile="sambamaster.db"
// first run with --init --passphrase xxxxxx --name xxxxx --password xxxx
// save startup passphrase in db
// create super administrator (adminlevel 65535)

// samba config
datadir="/samba/data"
dataowner="smbowner" // user and group
smbconf="/etc/samba/smb.conf"
smbpasswd="/usr/bin/smbpasswd"
smbd="/usr/sbin/smbd"
nmbd="/usr/sbin/nmbd"
winbind="/usr/sbin/winbind"
*/

/*
database:
passphrase: passphrase used for password encrypt
sharelist: name gid desc, name with 'user-' prefix is for personal share, one share for one system group
userlist: name uid password groups email cellphone adminlevel desc, password has encrypt by startup passphrase
userinfo: uid q1 q2 q3 a1 a2 a3
*/

/*
web pages:
1. no login welcome: show


*/

//
//
//
//
//
//
