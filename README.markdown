This repository integrates *Camlistore's* master branch with all changes
made by us here at *Toshnix Systems* that are not yet upstream in the
main *Camlistore* repo (see links below).

So the **master** branch on this repo is just an integration branch
of CLs that we are interested in testing together, which CLs we have
already submitted for review for integration into the main Camlistore
repo, but which are not yet upstream (At least, that is the basic
idea).

Some of the changes in this repo which are not yet upstream are:

* **FIFO** support for camput (uploading) and camget (restoring).
* **Socket** support for camput (uploading) and camget (restoring).
* Support for the use of **scrypt** and a passphrase for the *encrypt* blobserver.


For more information about *Camlistore*, see:

     http://camlistore.org/
     http://camlistore.org/docs/

Other useful files:

     BUILDING  how to compile it ("go run make.go")
     HACKING   how to do development and contribute

Mailing lists:

     http://camlistore.org/lists

Bugs and contributing:

     https://code.google.com/p/camlistore/issues/list
     http://camlistore.org/docs/contributing
