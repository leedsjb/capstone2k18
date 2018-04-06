import React from "react";
import { Flex } from "grid-styled";
import { Link } from "react-router-dom";

import Box from "../../components/Box";
import Heading from "../../components/Heading";
import Container from "../../components/Container";
import Logo from "../../components/Logo";
import Measure from "../../components/Measure";
import Button from "../../components/Button";
import Footer from "../../components/Footer";

const HomePage = () => {
    return (
        <div>
            <Box p={3}>
                <Flex justifyContent="space-between">
                    <Link to="/">
                        <Logo />
                    </Link>
                    <Link to="/missions">Sign in</Link>
                </Flex>
            </Box>
            <Container py={5}>
                <Measure>
                    <Heading is="h1" mt={0}>
                        Critical Communication for Airborne Medical Missions
                    </Heading>
                </Measure>
                <Measure mt={3}>
                    A mission-critical, high-availability application for
                    Airlift Northwest emergency flight nurses and pilots
                    transporting severely ill or injured patients by helicopter
                    to reach life-saving medical treatment.
                </Measure>
                <Button mt={3}>Call to action</Button>
            </Container>
            <Box bg="wireframe">
                <Container py={5}>
                    <Heading is="h2" fontSize={4} mt={0}>
                        Feature One
                    </Heading>
                    <Measure mt={2}>
                        Donec sed odio dui. Vivamus sagittis lacus vel augue
                        laoreet rutrum faucibus dolor auctor. Integer posuere
                        erat a ante venenatis dapibus posuere velit aliquet. Sed
                        posuere consectetur est at lobortis.
                    </Measure>

                    <Heading is="h2" fontSize={4} mt={4}>
                        Feature Two
                    </Heading>
                    <Measure mt={2}>
                        Maecenas sed diam eget risus varius blandit sit amet non
                        magna. Nulla vitae elit libero, a pharetra augue. Donec
                        ullamcorper nulla non metus auctor fringilla. Nulla
                        vitae elit libero, a pharetra augue.
                    </Measure>

                    <Heading is="h2" fontSize={4} mt={4}>
                        Feature Three
                    </Heading>
                    <Measure mt={2} mb={0}>
                        Maecenas sed diam eget risus varius blandit sit amet non
                        magna. Nulla vitae elit libero, a pharetra augue. Donec
                        ullamcorper nulla non metus auctor fringilla. Nulla
                        vitae elit libero, a pharetra augue.
                    </Measure>
                </Container>
            </Box>
            <Footer />
        </div>
    );
};

export default HomePage;
