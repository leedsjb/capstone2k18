import React from "react";
import { Flex } from "grid-styled";

import Box from "../Box";
import Container from "../Container";
import Image from "../Image";
import Span from "../Span";

import placeholder from "../../images/placeholder-image.jpg";

const Footer = () => {
    return (
        <div>
            <Box bg="wireframe" py={5}>
                <Container width={[1, 1 / 2, 1 / 2, 1 / 3]}>
                    <Span display="block" lineHeight={2} textAlign="center">
                        A University of Washington Information School Capstone
                        Project in collaboration with Airlift Northwest, a
                        University of Washington Medicine entity.
                    </Span>
                </Container>
            </Box>
            <Box px={5} py={4}>
                <Flex
                    flexDirection={["column", "row"]}
                    justifyContent="space-between"
                >
                    <Flex justifyContent="center" mb={2}>
                        <Box>
                            <Span lineHeight={2}>© 2018 Elevate</Span>
                        </Box>
                    </Flex>
                    <Flex justifyContent="space-between">
                        <Span lineHeight={2} mr={4}>
                            Contact us
                        </Span>
                        <Span lineHeight={2} mr={4}>
                            Privacy
                        </Span>
                        <Span lineHeight={2}>Terms</Span>
                    </Flex>
                </Flex>
            </Box>
        </div>
    );
};

export default Footer;
