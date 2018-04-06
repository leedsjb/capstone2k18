import React from "react";
import { Flex } from "grid-styled";
import { Link } from "react-router-dom";

import Box from "../../components/Box";
import Heading from "../../components/Heading";
import Container from "../../components/Container";
import Logo from "../../components/Logo";
import Measure from "../../components/Measure";
import Button from "../../components/Button";
import ButtonOutline from "../../components/ButtonOutline";
import Footer from "../../components/Footer";
import Divider from "../../components/Divider";
import Text from "../../components/Text";

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
                <Measure is="p" mt={3} mb={0}>
                    A mission-critical, high-availability application for
                    Airlift Northwest emergency flight nurses and pilots
                    transporting severely ill or injured patients by helicopter
                    to reach life-saving medical treatment.
                </Measure>
                <Button mt={3}>Get Started</Button>
            </Container>
            <Box bg="wireframe">
                <Container py={5}>
                    <Heading is="h2" fontSize={4} mt={0}>
                        Feature One
                    </Heading>
                    <Measure is="p" mt={2}>
                        Donec sed odio dui. Vivamus sagittis lacus vel augue
                        laoreet rutrum faucibus dolor auctor. Integer posuere
                        erat a ante venenatis dapibus posuere velit aliquet. Sed
                        posuere consectetur est at lobortis.
                    </Measure>

                    <Heading is="h2" fontSize={4} mt={4}>
                        Feature Two
                    </Heading>
                    <Measure is="p" mt={2}>
                        Maecenas sed diam eget risus varius blandit sit amet non
                        magna. Nulla vitae elit libero, a pharetra augue. Donec
                        ullamcorper nulla non metus auctor fringilla. Nulla
                        vitae elit libero, a pharetra augue.
                    </Measure>

                    <Heading is="h2" fontSize={4} mt={4}>
                        Feature Three
                    </Heading>
                    <Measure is="p" mt={2} mb={0}>
                        Praesent commodo cursus magna, vel scelerisque nisl
                        consectetur et. Morbi leo risus, porta ac consectetur
                        ac, vestibulum at eros. Lorem ipsum dolor sit amet,
                        consectetur adipiscing elit. Nulla vitae elit libero, a
                        pharetra augue.
                    </Measure>
                </Container>
            </Box>
            <Box py={5}>
                <Container>
                    <Heading is="h2" fontSize={4} mt={0}>
                        Nullam quis risus eget urna mollis ornare?<br />
                        <Text is="span" fontWeight="normal">
                            Maecenas sed diam eget risus varius.
                        </Text>
                    </Heading>
                    <Box mt={3}>
                        <Button>Get Started</Button>
                        <ButtonOutline ml={2}>Contact Airlift</ButtonOutline>
                    </Box>
                </Container>
            </Box>
            <Divider />
            <Footer />
        </div>
    );
};

export default HomePage;
