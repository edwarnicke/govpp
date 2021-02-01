# How the magic works #

This repo utilizes aggressively a number of multi-stage Docker tricks that are worth explaining, because they can be used
broadly

Broadly, we are doing three things in the Dockerfile, denoted by their named stage:

1. Building small vpp and vpp-dbg images, including the ability to patch
2. generating the govpp 'binapi' for the build vpp image, using the standard 'go:generate' idiom
3. Extracting the VPP_VERSION being used so we can utilize it to tag published docker images.

## Building small vpp images ##

Ultra small vpp images are built via three Docker stages:
1. vppbuild
2. vppinstall
3. vpp

The reason for these stages is to cauterize the bloat from teach activity.  

Building vpp means bloating an image
up with a bunch of build dependencies, build artifacts, etc.  Building vpp installation on top of that would 
lead to a multi-GB image, which is undesirable.  So we isolate that work in the 'vppbuild' stage.

Installing vpp means copying the resulting *.deb packages from the 'vppbuild' stage and installing them.  Unfortunately,
because you can't combine a COPY and RUN step, that results in image bloat from the *.deb files themselves, so we isolate
that in the 'vppinstall' image.

Finally, we utilize a trick to trim out the bloat when building the 'vpp' stage.

### The 'vppbuild' stage ###

https://github.com/edwarnicke/govpp/blob/f9c1af6818e193c701e8c34a6b17b5a3fafccceb/Dockerfile#L10-L20

is a fairly standard Ubuntu oriented build.  It results in a bunch of *.deb files.

### The 'vppinstall' stage ###

https://github.com/edwarnicke/govpp/blob/f9c1af6818e193c701e8c34a6b17b5a3fafccceb/Dockerfile#L22-L30

uses

https://github.com/edwarnicke/govpp/blob/f9c1af6818e193c701e8c34a6b17b5a3fafccceb/Dockerfile#L24

to copy in the indexes in ```/var/lib/apt/lists/``` that result from having run ```apt-get update``` in the 'vppbuild' stage,
thus avoiding the cost of redownloading them.

https://github.com/edwarnicke/govpp/blob/main/Dockerfile#L25

copies the *.deb files we wish to install from where they were built in 'vppbuild'

https://github.com/edwarnicke/govpp/blob/main/Dockerfile#L27

installs the *.deb files:

1. ```-f``` causes ```apt-get``` to install any missing dependencies.
2. ```-y``` causes ```apt-get``` to run in an unattended mode where the answer to the questions are ```y```
3, ``--no-install-recommends`` causes ```apt-get``` to only install required (rather than recommended) dependencies to keep the image size down

https://github.com/edwarnicke/govpp/blob/main/Dockerfile#L28

removes the apt indexes from ```apt-get update```

and

https://github.com/edwarnicke/govpp/blob/main/Dockerfile#L29

removes the *.deb files

There is one problem with this.  Because the image still has the layers from

https://github.com/edwarnicke/govpp/blob/main/Dockerfile#L24-L25

the 'vppinstall' stage is still going to be bloated by that amount.  We solve this in the 'vpp' stage

### The 'vpp' stage ###

The 'vpp' stage is our final lean runnable.  It uses a very simple but slick trick, which has a small caveat to it.

https://github.com/edwarnicke/govpp/blob/f9c1af6818e193c701e8c34a6b17b5a3fafccceb/Dockerfile#L31-L32

Simply starts from 'ubuntu:${UBUNTU_VERSION}' (thus reusing the layers form that standard image) and then
copies the entire '/' directory in from 'vppinstall' (which has removed the apt-get indexes and *.deb files).

Because docker COPY is generally smart enough to only copy in the *changed* files... the layer resulting from
https://github.com/edwarnicke/govpp/blob/f9c1af6818e193c701e8c34a6b17b5a3fafccceb/Dockerfile#L32
should only contain the deltas between the final result of 'vppinstall' and the starting point of 'ubuntu:${UBUNTU_VERSION}'

Resulting in something that looks like:
```
docker history ghcr.io/edwarnicke/govpp/vpp:v20.09                                                                                                                                                                                                                           ──(Sun,Jan31)─┘
IMAGE          CREATED          CREATED BY                                      SIZE      COMMENT
8b4ea0febd25   42 minutes ago   /bin/sh -c #(nop) COPY dir:b8f7abee062c48863…   96.2MB    
<missing>      10 days ago      /bin/sh -c #(nop)  CMD ["/bin/bash"]            0B        
<missing>      10 days ago      /bin/sh -c mkdir -p /run/systemd && echo 'do…   7B        
<missing>      10 days ago      /bin/sh -c [ -z "$(apt-get indextargets)" ]     0B        
<missing>      10 days ago      /bin/sh -c set -xe   && echo '#!/bin/sh' > /…   811B      
<missing>      10 days ago      /bin/sh -c #(nop) ADD file:2a90223d9f00d31e3…   72.9MB    
```

The caveat is, there is a [bug in docker](https://github.com/moby/moby/issues/21950) that will, depending on the [storage driver](https://docs.docker.com/storage/storagedriver/)
you are using, copy over all the files, resulting in a much larger image.  Fortunately, building in GitHub Actions does not
seem to hit this issue.  Unfortunately, building in [Docker for Mac](https://docs.docker.com/docker-for-mac/) does.
To fix the issue in Docker for Mac, follow [the instructions for setting Docker Engine options](https://docs.docker.com/docker-for-mac/#docker-engine)
and set your storage driver to 'overlay'
```json
{
  "storage-driver": "overlay"
}
```

## Building idiomatic go:generate for 'binapi' ##
In go, we generally generate code using a [//go:generate](https://blog.golang.org/generate) directive.

In the case of govpp, you need to use the 'binapi-generator' run against the json api files installed by the vpp debs in ```/usr/share/vpp/api/```.
We utilize Docker to make this easy by:

1. Building 'binapi-generator' as a Docker stage
2. Creating a 'gen' Docker stage that copies in ```/usr/share/vpp/api/``` from our 'vpp' stage and run's binapi, outputting to /gen
3. Add a ```//go:generate``` line to [gen.go](https://github.com/edwarnicke/govpp/blob/f9c1af6818e193c701e8c34a6b17b5a3fafccceb/gen.go) to run the 'gen' stage to generate the code.

### The 'bin-apigenerator' stage ###

https://github.com/edwarnicke/govpp/blob/main/Dockerfile#L41-46

Is a pretty standard 'go get to build' stage for 'binapi-generator'

### The 'gen' stage ###
https://github.com/edwarnicke/govpp/blob/main/Dockerfile#L48-53

Actually performs the generation if 'docker run'

https://github.com/edwarnicke/govpp/blob/main/Dockerfile#L49

copies in the '/usr/share/vpp/api/*' json api files from the 'vpp' stage

https://github.com/edwarnicke/govpp/blob/main/Dockerfile#L50

copies in the 'binapi-generator' from the 'binapi-generator' stage

https://github.com/edwarnicke/govpp/blob/f9c1af6818e193c701e8c34a6b17b5a3fafccceb/Dockerfile#L51

copies in the VPP_VERSION we had [stashed](https://github.com/edwarnicke/govpp/blob/f9c1af6818e193c701e8c34a6b17b5a3fafccceb/Dockerfile#L20) in the 'vppbuild' stage.

https://github.com/edwarnicke/govpp/blob/main/Dockerfile#L52-L53

sets the Workdir to /gen and runs the binapi-generator.

```VPP_VERSION=$(cat /VPP_VERSION)``` - sets the VPP_VERSION env binapi-generator needs from the stashed VPP_VERSION value

```${PKGPREFIX+-import-prefix ${PKGPREFIX}}``` - will output ```-import-prefix ${PKGPREFIX}``` if the PKGPREFIX env variable is set, and nothing otherwise.

### The ```//go:generate``` directive ###

https://github.com/edwarnicke/govpp/blob/f9c1af6818e193c701e8c34a6b17b5a3fafccceb/gen.go#L19

uses 'docker run' to run the 'gen' stage and generate the code

1. ```bash bash -c "docker run ... "``` is used to give us a shell to work with (because we are doing a lot of magic here)
2. ```-e PKGPREFIX=$(go list)/binapi``` - sets the env variable ```PKGPREFIX``` in the docker container to the value of
   ```$(go list)/binapi```.  ```$(go list)``` is the value of the package in which the gen.go file resides.
3. ```-v $(go list -f '{{ .Dir }}'):/gen``` mounts ```$(go list -f '{{ .Dir }}')``` from the host into ```/gen/``` in the container.
    1. ```$(go list -f '{{ .Dir }}')``` outputs the directory the module containing gen.go is in.
4. ```$(docker build . -q --build-arg GOVPP_VERSION=$(go list -m -f '{{ .Version }}'  git.fd.io/govpp.git))```
    1. ```-q``` - 'quiet' - outputs the id of the resulting image built
    1. ```--build-arg GOVPP_VERSION=$(go list -m -f '{{ .Version }}'  git.fd.io/govpp.git))``` - sets the build-arg GOVPP_VERSION 
    1. ```$(go list -m -f '{{ .Version }}'  git.fd.io/govpp.git)``` -  outputs the version of git.fd.io/govpp.git in the go.mod file.

### Extracting the VPP_VERSION being used in the Dockerfile ###

In https://github.com/edwarnicke/govpp/blob/main/.github/workflows/ci.yaml we want to be able to tag and push images based
on the VPP_VERSION.  Because we only wish to specify the VPP_VERSION *once* in the Dockerfile, we have a job:

https://github.com/edwarnicke/govpp/blob/f9c1af6818e193c701e8c34a6b17b5a3fafccceb/Dockerfile#L5-L8

that can be used to extract that version:
https://github.com/edwarnicke/govpp/blob/f9c1af6818e193c701e8c34a6b17b5a3fafccceb/.github/workflows/ci.yaml#L92

