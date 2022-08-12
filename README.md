### taskMe!

taskMe! is a simple simple todo Desktop app written in Go with [Fyne](https://developer.fyne.io/).

##### An illustration of v1.0.0
![alt text](https://github.com/petrostrak/task-me/blob/main/taskMe.png)

##### An illustration of v2.0.0 / v3.0.0
![alt text](https://github.com/petrostrak/task-me/blob/main/taskMev2.png)

##### An illustration of v4.0.0
![alt text](https://github.com/petrostrak/task-me/blob/main/taskMev3.png)

##### To build taskMe! from source
    $ fyne package -appVersion 1.0.0 -appBuild 1 -name taskMe -release

##### To install taskMe! on debian:

* Extract taskMe.tar.xz
    
    `mkdir taskME && tar -xvf taskME.tar.xz -C taskME`

* Mount and run make file:

    `cd taskME && sudo make install`
