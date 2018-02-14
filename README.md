# capstone2k18
University of Washington Informatics Capstone 2018

## Development Flow
Feature work should be done on branches named for the feature that is being developed. Once completed, create a pull request to merge
your feature into the development branch for internal testing. To deploy to production the development branch will be merged to master. 

To work with Google Cloud Platform (GCP) you may wish to install the Google Cloud SDK. The SDK allows you to interact with GCP from your
command line. To install for macOS visit Google's [Quickstart Article](https://cloud.google.com/sdk/docs/quickstart-macos) to download the SDK.
Once the initial install is complete be sure to:
    1.  Run `gcloud init` to sign in to your Google account and choose a project to work with.
    2.  Install the kubectl component to interact with Kubernetes. The component can be installed by running `gcloud components install kubectl` 
        from your home directory. 

## Resources
Articles/documentation that together explain our UI rationale. Meant to be read in order. Tutorials are optional.


### Component System

#### Folder Structure
[Components vs. containers](https://medium.com/@dan_abramov/smart-and-dumb-components-7ca2f9a7c7d0)

[Scaling React applications (only read the folder part)](https://www.smashingmagazine.com/2016/09/how-to-scale-react-applications/)

#### Styled Components
[Documentation](https://www.styled-components.com/docs)

[Rationale (optional)](https://www.youtube.com/watch?time_continue=89&v=bIK2NwoK9xk)

#### Rebass
[Documentation](http://jxnblk.com/rebass/)


### Design System
[The 8-Point Grid](https://spec.fm/specifics/8-pt-grid)


### Working With Data

#### Redux:
[Documentation](https://redux.js.org/)

[Basic Redux tutorial](https://egghead.io/courses/getting-started-with-redux)

[Advanced Redux tutorial](https://egghead.io/courses/building-react-applications-with-idiomatic-redux)

[react-redux](https://github.com/reactjs/react-redux)

[redux-thunk](https://github.com/gaearon/redux-thunk)

[redux-form](https://redux-form.com)

[reselect](https://github.com/reactjs/reselect)

