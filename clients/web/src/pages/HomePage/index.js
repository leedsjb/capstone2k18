import React from "react";
import { Flex } from "grid-styled";

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
                    <Logo />
                    <span>Sign in</span>
                </Flex>
            </Box>
            <Container py={5}>
                <Heading is="h1" mt={0}>
                    AirliftNW Elevate
                </Heading>
                <Measure mt={3}>
                    Nullam quis risus eget urna mollis ornare vel eu leo. Nullam
                    quis risus eget urna mollis ornare vel eu leo. Donec id elit
                    non mi porta gravida at eget metus. Maecenas sed diam eget
                    risus varius blandit sit amet non magna. Sed posuere
                    consectetur est at lobortis.
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
