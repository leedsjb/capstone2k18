import React from "react";

import Text from "../Text";
import Container from "../Container";
import Box from "../Box";

const Footer = () => {
    return (
        <Box py={5}>
            <Container>
                <Text my={0}>
                    Donec id elit non mi porta gravida at eget metus
                </Text>
            </Container>
        </Box>
    );
};

export default Footer;
