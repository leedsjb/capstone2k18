import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { Container, Heading, Text } from "rebass";

import StyledMeasure from "../../components/StyledMeasure";

class AircraftsPage extends Component {
    render() {
        return (
            <div>
                <Helmet>
                    <title>Aircrafts</title>
                </Helmet>
                <Container>
                    <StyledMeasure>
                        <Heading is="h1">Aircrafts</Heading>
                        <Text>
                            Aliquam dignissim, est nec luctus dapibus, velit mi
                            tristique nulla, eget hendrerit velit tortor vitae
                            risus. Nulla mollis ex ut lorem finibus ultricies.
                        </Text>
                    </StyledMeasure>
                </Container>
            </div>
        );
    }
}

export default AircraftsPage;
