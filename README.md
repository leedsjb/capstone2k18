# Airlift Northwest Elevate: Delivering lifesaving care, faster.
[Live](http://test.elevate.airliftnw.org/)

Elevate is a mission-critical, high-availability application for Airlift Northwest emergency flight nurses and pilots in Washington and Alaska. These professionals transport severely ill or injured patients by airplane and helicopter to reach lifesaving medical treatment. We bridge the communication gap between the dispatch center on the ground and flight crews in the air by displaying real-time mission status. Elevate ensures flight crews arrive at the right location with the equipment and information needed to immediately administer lifesaving patient care.

## Features

### Aircraft

* Track aircraft. For every Airlift aircraft, see its current location and status, as well as any additional information such as patient information
* Search for specific aircraft
* Filter aircraft by status and/or category

### People

* Look up contact information of Airlift Northwest employees
* Browse organization by people or groups
* Search for specific people or groups

### Resources

* Access important Airlift Northwest resources, such as Ninth Brain LMS, in two clicks

## Project Structure

* clients/web: contains the web client
* clients/websocket: contains a client for WebSocket testing
* deployment: contains deployment setup
* servers/elevate: 
* servers/mysql; 

## Rationale

### Frontend

* Why React? We chose to use React because its virtual dom would help ensure our app is performant and its component structure makes projects easy to manage.
* Why styled-components? Ideally, a component based project enables the developer to move component folders around without having to do anything else. Default React styling (regular CSS style sheets) does not accommodate for this behavior, making it harder to manage components in one place. styled-components, CSS-in-JS, makes components independent entities by coupling markup and style. Furthermore, styled-components comes with several other benefits such as theming. Theming makes it possible to guarantee consistency throughout large web apps.
* Why Redux? While React does not have to be used with Redux, we chose to use Redux since without it we would constantly have to lift state that is relevant to more than one component, such as aircraft information, up. Redux also makes it easy to inform components about WebSocket updates.
* Why Mapbox, and not other solutions such as Leaflet? We chose Mapbox after talking to Dr. Stearns and expressing our concerns over the performance of maps in browsers. He recommended it because it is powered by OpenGL and thus ideal when performance is a must.

### Backend

## Contact

J. Benjamin Leeds
leeds@uw.edu

Tiffany Chen
tzchen@uw.edu

Jessica Basa
jdbasa24@uw.edu

Vincent van der Meulen
meulen@uw.edu
