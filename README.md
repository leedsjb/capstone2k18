# Airlift Northwest Elevate: Delivering lifesaving care, faster.
[Live](https://test.elevate.airliftnw.org/)

Elevate is a mission-critical, high-availability application for Airlift Northwest emergency flight nurses and pilots in Washington and Alaska. These professionals transport severely ill or injured patients by airplane and helicopter to reach lifesaving medical treatment. We bridge the communication gap between the dispatch center on the ground and flight crews in the air by displaying real-time mission status. Elevate ensures flight crews arrive at the right location with the equipment and information needed to immediately administer lifesaving patient care.

## Features

### Aircraft

* Track aircraft. For every Airlift aircraft, see its current location and status, as well as any additional information such as patient information
* Search for specific aircraft
* Filter aircraft by status and/or category

### Missions
* As a pilot or flight nurse, receive notifications of new missions assigned to me and provide a deep link to view more detailed mission information in Elevate

### People

* Look up contact information of Airlift Northwest employees
* Browse organization by people or groups
* Search for specific people or groups

### Resources

* Access important Airlift Northwest resources, such as Ninth Brain LMS, in two clicks

## Project Structure

* clients/web: contains the React web client
* clients/websocket: contains a client for WebSocket testing
* deployment: contains deployment configuration files (Kubernetes YAML and associated scripts)
* servers/elevate: the API server that handles client requests and ingests mission data from the ALNW dispatch center 
* servers/mysql: the configuration scripts for the MySQL server including: DDL, Stored Procedures, and test data load scripts

## Rationale

### Frontend

* Why React? We chose to use React because its virtual dom would help ensure our app is performant and its component structure makes projects easy to manage.
* Why styled-components? Ideally, a component based project enables the developer to move component folders around without having to do anything else. Default React styling (regular CSS style sheets) does not accommodate for this behavior, making it harder to manage components in one place. styled-components, CSS-in-JS, makes components independent entities by coupling markup and style. Furthermore, styled-components comes with several other benefits such as theming. Theming makes it possible to guarantee consistency throughout large web apps.
* Why Redux? While React does not have to be used with Redux, we chose to use Redux since without it we would constantly have to lift state that is relevant to more than one component, such as aircraft information, up. Redux also makes it easy to inform components about WebSocket updates.
* Why Mapbox, and not other solutions such as Leaflet? We chose Mapbox after talking to Dr. Stearns and expressing our concerns over the performance of maps in browsers. He recommended it because it is powered by OpenGL and thus ideal when performance is a must.

### Backend

* Why Go? Google Golang was designed at Google with building performant API servers in mind. Golang's use of Go Routines allow requests to an API server to be processed in parallel on multiple processors to handle high traffic loads efficiently. Combined with cluster autoscaling using Kubernetes we were able to design an application that can handle many more requests than all Airlift Northwest employee users (200+) can make simultaneously. 

* Why Kubernetes (K8S)?  We wanted to ensure that bringing Airlift Northwest the most up-to-date technology would be both sustainable and highly reliable. Kubernetes allows us as application designers to specify the "desired state" of the application architecture (container images, # of virtual machines (nodes), network ingress configuration, TLS and application secrets) as code using a YAML file. Once the file is applied to a K8S cluster a master node actively keeps the cluster in the specified state automatically. This allows Elevate to stay up and running with minimal developer intervention required. 

* Why Google Cloud? Airlift's lack of cloud storage limits the reliability and durability of their data. By using Google Cloud, we were able to avoid infrastructure costs, bypass the acquisition and installation of compute resources that would otherwise have been necessary, and structure our application in a scalable and highly available way.

* Why MySQL? MySQL is an industry standard, well supported by development tools and application frameworks. This allows for flexibility on our end to tailor it to our needs. Additionally MySQL is open source helping to keep application costs low for our nonprofit sponsor. 

## Contact

J. Benjamin Leeds
leeds@uw.edu

Tiffany Chen
tzchen@uw.edu

Jessica Basa
jdbasa24@uw.edu

Vincent van der Meulen
meulen@uw.edu
