import React from "react";
import { Flex } from "grid-styled";

import Box from "../Box";
import Container from "../Container";
import Image from "../Image";
import Span from "../Span";

import collab from "../../images/collab.svg";
import symbol from "../../images/symbol.svg";

const Footer = () => {
    return (
        <div>
            <Box py={8}>
                <Container width={[7 / 8, 1 / 2]}>
                    <Flex flexDirection="column" justifyContent="center">
                        <Box mb={4}>
                            <Flex justifyContent="center">
                                <Image
                                    src={collab}
                                    alt="AirliftNW Elevate and UW Information School logos"
                                />
                            </Flex>
                        </Box>
                        <Box width={1}>
                            <Span
                                color="#515766"
                                display="block"
                                fontWeight="100"
                                lineHeight={2}
                                textAlign="center"
                            >
                                A University of Washington Information School
                                Capstone Project in collaboration with Airlift
                                Northwest, a University of Washington Medicine
                                entity.
                            </Span>
                        </Box>
                    </Flex>
                </Container>
            </Box>
            <Box pb={4} pt={[0, 0, 4]} px={5}>
                <Flex
                    flexDirection={["column", "row"]}
                    justifyContent="space-between"
                >
                    <Flex justifyContent="center" mb={2}>
                        <Box mr={2} mt={1}>
                            <Image
                                src={symbol}
                                alt="AirliftNW Elevate"
                                w={40}
                            />
                        </Box>
                        <Box>
                            <Span fontSize={1} lineHeight={2}>
                                © 2018 Elevate
                            </Span>
                        </Box>
                    </Flex>
                    <Flex justifyContent="space-between">
                        <a
                            href="https://www.uwmedicine.org/airlift-nw/contact-us"
                            rel="noopener noreferrer"
                            target="_blank"
                        >
                            <Span
                                fontSize={1}
                                fontWeight="bold"
                                lineHeight={2}
                                mr={4}
                            >
                                Contact us
                            </Span>
                        </a>
                        <a
                            href="https://www.uwmedicine.org/about/compliance/privacy"
                            rel="noopener noreferrer"
                            target="_blank"
                        >
                            <Span
                                fontSize={1}
                                fontWeight="bold"
                                lineHeight={2}
                                mr={4}
                            >
                                Privacy
                            </Span>
                        </a>
                        <a
                            href="https://www.washington.edu/online/terms/"
                            rel="noopener noreferrer"
                            target="_blank"
                        >
                            <Span fontSize={1} fontWeight="bold" lineHeight={2}>
                                Terms
                            </Span>
                        </a>
                    </Flex>
                </Flex>
            </Box>
        </div>
    );
};

export default Footer;
