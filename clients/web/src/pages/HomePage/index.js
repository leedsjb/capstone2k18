import React from "react";
import { Flex } from "grid-styled";

import Box from "../../components/Box";
import Heading from "../../components/Heading";
import Container from "../../components/Container";
import Logo from "../../components/Logo";

const HomePage = () => {
    return (
        <div>
            <Box p={3}>
                <Flex justifyContent="space-between">
                    <Logo />
                    <span>Sign in</span>
                </Flex>
            </Box>
            <Container>
                <Heading is="h1">AirliftNW Elevate</Heading>
                <div>Home Page</div>
            </Container>
        </div>
    );
};

export default HomePage;
