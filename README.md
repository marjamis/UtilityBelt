#Testing Platform for connections to multiple services when used with docker

Each service will have the minimum functionality of:
Add
Display
Delete
and maybe more based on the serivce.


ToDo:
- When AMZN has docker 1.10 change to compose version 2.
- Overlay network with consul
- HAProxy infront of applications
- DNS Records for HAProxy
- Have the entire environment creation automated
- add password to the redis connection
- Integrate with splunk? caching? like varnish?
- troposhere for CFN generation if I go down that path
- mysql with secure connections with about x509 and ssl guide - https://dev.mysql.com/doc/refman/5.0/en/
- if adding swarm functionality have the swarm manager HA configured, at least 3 nodes(with normal HA stuffs like separate AZs), HA for consul as well, maybe a template or docker machine script to set all the above up outside of normal processes? - seperate consul for application/overlay networks than for swarm related things?? maybe overkill but make it a point somehwere
- integrate a CI/CD set up into this a bit for testing? jenkins container which monitors the code base to build new image of my application to replace and run on swarm/ecs? Re-read:
https://www.docker.com/sites/default/files/UseCase/RA_CI%20with%20Docker_08.25.2015.pdf?iesrc=rcmd&astid=e060df8b-1bce-4897-a31a-9ce6d510081d&at=1&rcmd_source=BAR&req_id=70128954-5ee8-45e4-b6e0-4316a6f8bae1
when I get here in the process
- Add in my delay request to this for general testing
- Add Kibana
- Add to tests: <get some basic ones and grow> language lint unit testing/Service/UI pipeline lambda
- CODE PIPELINE FOR DEPLOYING/TESTING
- Add monitoring configurations for every step of infra/code such as with datadog
- scrum software, devops stuff -- use tuleap.org
- try with gitlabd
- keep 12factor.net in mind
- build an rpm?
- jenkins best practices
- git hook for golang vet command and golang/lint prior to commits to ensure its all good
