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
            <Container py={12}>
                <Flex flexDirection="column" justifyContent="center">
                    <Box>
                        <Flex justifyContent="center">
                            <Image
                                src={collab}
                                alt="AirliftNW Elevate and UW Information School logos"
                            />
                        </Flex>
                    </Box>
                    <Box w={1} maxWidth="30em" mx="auto">
                        <Span
                            color="#515766"
                            display="block"
                            fontWeight="100"
                            lineHeight={2}
                            textAlign="center"
                            mt={8}
                        >
                            A University of Washington Information School
                            Capstone Project in collaboration with Airlift
                            Northwest, a University of Washington Medicine
                            entity.
                        </Span>
                    </Box>
                </Flex>
                <Flex
                    flexDirection={["column", "row"]}
                    justifyContent="space-between"
                    mt={18}
                >
                    <Flex justifyContent="center">
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
            </Container>
        </div>
    );
};

export default Footer;
