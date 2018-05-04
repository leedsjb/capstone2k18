import React from "react";
import { Flex } from "grid-styled";
import { Link } from "react-router-dom";

import Box from "../../components/Box";
import Button from "../../components/Button";
import ButtonOutline from "../../components/ButtonOutline";
import Container from "../../components/Container";
import Divider from "../../components/Divider";
import Footer from "../../components/Footer";
import Heading from "../../components/Heading";
import Image from "../../components/Image";
import Logo from "../../components/Logo";
import Measure from "../../components/Measure";
import RedBox from "../../components/RedBox";
import Span from "../../components/Span";
import Text from "../../components/Text";

import placeholder from "../../images/placeholder-image.jpg";
import phone from "../../images/phone.png";
import logo from "../../images/logo.svg";
import second from "../../images/second.png";

const HomePage = () => {
    return (
        <div>
            <Box bg="#F7F9FC">
                <Container>
                    <Flex
                        justifyContent="space-between"
                        alignItems="center"
                        pt={3}
                    >
                        <Link to="/">
                            <Image src={logo} alt="AirliftNW Elevate" w={183} />
                        </Link>
                        <div>
                            <Button>
                                <Link to="/signin">Sign in</Link>
                            </Button>
                        </div>
                    </Flex>

                    <Flex
                        flexWrap={["wrap", "wrap", "nowrap"]}
                        justifyContent="space-between"
                        alignItems="center"
                    >
                        <Box width={[1, 1, 2 / 3]}>
                            <Heading is="h1" mt={0} fontSize={6}>
                                Critical Communication for Airborne Medical
                                Missions
                            </Heading>
                            <Measure mt={3} mb={0}>
                                Elevate is a mission-critical, high-availability
                                application for Airlift Northwest emergency
                                flight nurses and pilots transporting severely
                                ill or injured patients by helicopter to reach
                                life-saving medical treatment.
                            </Measure>
                            <Button mt={3}>Get Started</Button>
                        </Box>
                        <Box my={5} width={[1, 1, 1 / 3]}>
                            <Image src={phone} width={1} />
                        </Box>
                    </Flex>
                </Container>
            </Box>

            <RedBox>
                <Container>
                    <Box py={6}>
                        <Flex
                            flexWrap={["wrap", "wrap", "nowrap"]}
                            justifyContent="space-between"
                            alignItems="center"
                        >
                            <Box maxWidth={488}>
                                <Image src={second} width={1} />
                            </Box>
                            <Box>
                                <Heading
                                    is="h2"
                                    fontSize={4}
                                    pt={4}
                                    color="white"
                                >
                                    Because every second counts
                                </Heading>
                                <Measure mt={2} color="white">
                                    Elevate bridges the communication gap
                                    between the dispatch center on the ground
                                    and flight crews in the air by displaying
                                    patient, aircraft, and aircrew status. This
                                    ensures that aircraft and aircrews quickly
                                    arrive at the right location with the
                                    information needed to immediately administer
                                    life-saving patient care.
                                </Measure>
                            </Box>
                        </Flex>
                    </Box>
                </Container>
            </RedBox>

            <Box>
                <Container px={5} py={5}>
                    <Heading is="h2" fontSize={5} mb={4} mt={2}>
                        Elevating health care technology
                    </Heading>
                    <Flex flexWrap={["wrap", "wrap", "nowrap"]}>
                        <Box mr={4}>
                            <Image src={placeholder} width={1 / 6} />
                            <Heading is="h4" fontSize={3} mt={2}>
                                Use anywhere
                            </Heading>
                            <Measure mt={2}>
                                Fusce dapibus, tellus ac cursus commodo, tortor
                                mauris condimentum nibh, ut fermentum massa
                                justo sit.
                            </Measure>
                        </Box>

                        <Box mr={4}>
                            <Image src={placeholder} width={1 / 6} />
                            <Heading is="h4" fontSize={3} mt={2}>
                                HIPAA Compliant
                            </Heading>
                            <Measure mt={2}>
                                Fusce dapibus, tellus ac cursus commodo, tortor
                                mauris condimentum nibh, ut fermentum massa
                                justo sit amet.
                            </Measure>
                        </Box>

                        <Box>
                            <Image src={placeholder} width={1 / 6} />
                            <Heading is="h4" fontSize={3} mt={2}>
                                Works offline
                            </Heading>
                            <Measure mt={2}>
                                Fusce dapibus, tellus ac cursus commodo, tortor
                                mauris condimentum nibh, ut fermentum massa
                                justo sit amet.
                            </Measure>
                        </Box>
                    </Flex>
                </Container>
            </Box>

            <RedBox>
                <Container px={5} py={5}>
                    <Box>
                        <Flex
                            flexWrap={["wrap", "wrap", "nowrap"]}
                            flexDirection={[
                                "column-reverse",
                                "column-reverse",
                                "row"
                            ]}
                            alignItems="center"
                        >
                            <Box mr={[0, 0, 5]}>
                                <Heading
                                    is="h2"
                                    fontSize={4}
                                    pt={5}
                                    color="white"
                                >
                                    Track missions
                                </Heading>
                                <Measure mt={2} color="white">
                                    Duis mollis, est non commodo luctus, nisis
                                    erat porttitor ligula, eget lacinia odio sem
                                    nec elit. Nulla vitae elit libero, a
                                    pharetra augue. Vistibulum id lignula port
                                    fellis euismod semper.
                                </Measure>
                            </Box>
                            <Box maxWidth={488}>
                                <Image src={placeholder} width={1} />
                            </Box>
                        </Flex>
                    </Box>
                </Container>
            </RedBox>

            <Box bg="#F7F9FC">
                <Container>
                    <Flex
                        flexWrap={["wrap", "wrap", "nowrap"]}
                        alignItems="center"
                    >
                        <Box maxWidth={488}>
                            <Image src={placeholder} width={1} />
                        </Box>
                        <Box>
                            <Heading is="h2" fontSize={4} pt={5}>
                                Get notified
                            </Heading>
                            <Measure mt={2}>
                                Duis mollis, est non commodo luctus, nisis erat
                                porttitor ligula, eget lacinia odio sem nec
                                elit. Nulla vitae elit libero, a pharetra augue.
                                Vistibulum id lignula port fellis euismod
                                semper.
                            </Measure>
                        </Box>
                    </Flex>
                </Container>
            </Box>

            <RedBox>
                <Container py={5}>
                    <Box>
                        <Flex
                            flexWrap={["wrap", "wrap", "nowrap"]}
                            flexDirection={[
                                "column-reverse",
                                "column-reverse",
                                "row"
                            ]}
                            alignItems="center"
                        >
                            <Box mr={[0, 0, 5]}>
                                <Heading
                                    is="h2"
                                    fontSize={4}
                                    pt={5}
                                    color="white"
                                >
                                    Find the right person
                                </Heading>
                                <Measure mt={2} color="white">
                                    > Duis mollis, est non commodo luctus, nisis
                                    erat porttitor ligula, eget lacinia odio sem
                                    nec elit. Nulla vitae elit libero, a
                                    pharetra augue. Vistibulum id lignula port
                                    fellis euismod semper.
                                </Measure>
                            </Box>
                            <Box width={1}>
                                <Image src={placeholder} width={1} />
                            </Box>
                        </Flex>
                    </Box>
                </Container>
            </RedBox>

            <Box>
                <Container py={5}>
                    <Box>
                        <Flex
                            flexWrap={["wrap", "wrap", "nowrap"]}
                            alignItems="center"
                        >
                            <Box mr={[0, 0, 5]}>
                                <Image src={placeholder} width={1} />
                            </Box>
                            <Box>
                                <Heading is="h2" fontSize={4} pt={4}>
                                    In partnership with Airlift Northwest
                                </Heading>
                                <Measure mt={2}>
                                    Curabitus blandit tempus porttitor. Donec
                                    ullamcorper nulla non metus auctor
                                    fringilla. Nullam quis risus get urna nollis
                                    ornare vel eu leo. Maecenas fancibus mollis
                                    interdum.
                                </Measure>
                                <Button>Get to know Airlift</Button>
                            </Box>
                        </Flex>
                    </Box>
                </Container>
            </Box>

            <RedBox>
                <Container px={5} py={5}>
                    <Box>
                        <Heading is="h2" fontSize={4} pt={2} mb={3}>
                            We elevate
                        </Heading>
                        <Flex flexWrap={["wrap", "wrap", "nowrap"]}>
                            <Box mb={3} mr={[0, 0, 5]}>
                                <Image src={placeholder} width={1} />
                                <Heading
                                    is="h5"
                                    fontSize={2}
                                    textAlign="center"
                                    my={2}
                                >
                                    Benjamin Leeds
                                </Heading>
                                <Span display="block" textAlign="center">
                                    TPM and Developer
                                </Span>
                            </Box>
                            <Box mb={3} mr={[0, 0, 5]}>
                                <Image src={placeholder} width={1} />
                                <Heading
                                    is="h5"
                                    fontSize={2}
                                    textAlign="center"
                                    my={2}
                                >
                                    Jessica Basa
                                </Heading>
                                <Span display="block" textAlign="center">
                                    Developer
                                </Span>
                            </Box>
                            <Box mb={3} mr={[0, 0, 5]}>
                                <Image src={placeholder} width={1} />
                                <Heading
                                    is="h5"
                                    fontSize={2}
                                    textAlign="center"
                                    my={2}
                                >
                                    Tiffany Chen
                                </Heading>
                                <Span display="block" textAlign="center">
                                    Designer and Developer
                                </Span>
                            </Box>
                            <Box mb={3}>
                                <Image src={placeholder} width={1} />
                                <Heading
                                    is="h5"
                                    fontSize={2}
                                    textAlign="center"
                                    my={2}
                                >
                                    Vincent van der Meulen
                                </Heading>
                                <Span display="block" textAlign="center">
                                    Designer and Developer
                                </Span>
                            </Box>
                        </Flex>
                    </Box>
                </Container>
            </RedBox>

            <Box mt={5}>
                <Container width={1 / 2}>
                    <Image src={placeholder} width={1} />
                </Container>
            </Box>
            <Flex flex="0 1 100%" flexWrap="wrap" justifyContent="center">
                <Box>
                    <Container px={5} py={5}>
                        <Heading is="h2" fontSize={4} textAlign="center" mb={2}>
                            Ready for takeoff?
                        </Heading>
                        <Span fontWeight="normal">
                            Elevate is available to Airlift Northwest employees
                        </Span>
                        <Box mt={3}>
                            <Button>Get Started</Button>
                            <ButtonOutline ml={2}>
                                Contact Airlift
                            </ButtonOutline>
                        </Box>
                    </Container>
                </Box>
            </Flex>
            <Divider />
            <Footer />
        </div>
    );
};

export default HomePage;
