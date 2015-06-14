docker run -v $(pwd):/isorender -dit --name=isorender -p3000:3000 node /bin/bash
docker run -v $(pwd):/isorender -dit --name=go -p 5000:5000 golang /bin/bash
docker exec -it go /bin/bash
