import React, { Component } from "react";
import { Helmet } from "react-helmet";

import Text from "../../components/Text";

class AircraftsPage extends Component {
    render() {
        return (
            <div>
                <Helmet>
                    <title>Aircrafts</title>
                </Helmet>
                <div>Aircrafts page</div>
                <Text>
                    Morbi leo risus, porta ac consectetur ac, vestibulum at
                    eros. Maecenas sed diam eget risus varius blandit sit amet
                    non magna. Nullam id dolor id nibh ultricies vehicula ut id
                    elit. Vivamus sagittis lacus vel augue laoreet rutrum
                    faucibus dolor auctor. Praesent commodo cursus magna, vel
                    scelerisque nisl consectetur et.
                </Text>
            </div>
        );
    }
}

export default AircraftsPage;
