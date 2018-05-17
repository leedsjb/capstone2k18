import React from "react";
import { Flex } from "grid-styled";
import { Link } from "react-router-dom";

import Box from "../../components/Box";
import Button from "../../components/Button";
import ButtonOutline from "../../components/ButtonOutline";
import Container from "../../components/Container";
import Footer from "../../components/Footer";
import Heading from "../../components/Heading";
import Image from "../../components/Image";
import Measure from "../../components/Measure";
import RedBox from "../../components/RedBox";
import Span from "../../components/Span";
import Icon from "../../components/Icon";

import benjamin from "../../images/benjamin.png";
import heli from "../../images/heli.png";
import logo from "../../images/logo.svg";
import search from "../../images/search.png";
import jessica from "../../images/jessica.png";
import partner from "../../images/partner.png";
import phone from "../../images/phone.png";
import second from "../../images/second.png";
import tiffany from "../../images/tiffany.png";
import vincent from "../../images/vincent.png";
import notification from "../../images/notification.svg";
import trackMissions from "../../images/trackMissions.png";

const HomePage = () => {
    return (
        <div>
            <Box bg="#F7F9FC">
                <Container>
                    <Flex justifyContent="space-between" pt={4}>
                        <Link to="/">
                            <Image src={logo} alt="AirliftNW Elevate" w={144} />
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
                            <Link to="/signin">
                                <Button mt={3}>Get Started</Button>
                            </Link>
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
                                <Heading is="h2" fontSize={4} color="white">
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
                    <Heading is="h2" fontSize={5} mb={4}>
                        Elevating health care technology
                    </Heading>
                    <Flex flexWrap={["wrap", "wrap", "nowrap"]} mt={5}>
                        <Box mr={4}>
                            <Icon glyph="devices" size={48} />
                            <Heading is="h4" fontSize={3} mt={3}>
                                Use anywhere
                            </Heading>
                            <Measure mt={2}>
                                Elevate is available on any device with a web
                                browser.
                            </Measure>
                        </Box>

                        <Box mr={4}>
                            <Icon glyph="checkShield" size={48} />
                            <Heading is="h4" fontSize={3} mt={3}>
                                HIPAA Compliant
                            </Heading>
                            <Measure mt={2}>
                                We comply with HIPAA and do not store or show
                                any PHI.
                            </Measure>
                        </Box>

                        <Box>
                            <Icon glyph="wifiCheck" size={48} />
                            <Heading is="h4" fontSize={3} mt={3}>
                                Works offline
                            </Heading>
                            <Measure mt={2}>
                                If you lose your internet connnection, the most
                                recently updated information will still be
                                available to you.
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
                                    Know what is going on at all times and
                                    follow missions on a map as they are
                                    happening.
                                </Measure>
                            </Box>
                            <Box maxWidth={488}>
                                <Image src={trackMissions} width={1} />
                            </Box>
                        </Flex>
                    </Box>
                </Container>
            </RedBox>

            <Box bg="#F7F9FC" py={5}>
                <Container>
                    <Flex
                        flexWrap={["wrap", "wrap", "nowrap"]}
                        alignItems="center"
                    >
                        <Box maxWidth={488}>
                            <Image src={notification} width={1} />
                        </Box>
                        <Box>
                            <Heading is="h2" fontSize={4}>
                                Get notified
                            </Heading>
                            <Measure mt={2}>
                                Receive a text notification as soon as you are
                                assigned to a new mission.
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
                                <Heading is="h2" fontSize={4} color="white">
                                    Find the right person
                                </Heading>
                                <Measure mt={2} color="white">
                                    Find and contact the person or group you’re
                                    looking for in seconds.
                                </Measure>
                            </Box>
                            <Box maxWidth={488}>
                                <Image src={search} width={1} />
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
                                <Image src={partner} width={1} />
                            </Box>
                            <Box>
                                <Heading is="h2" fontSize={4}>
                                    About Airlift Northwest
                                </Heading>
                                <Measure mt={2}>
                                    Elevate was built together with Airlift
                                    Northwest, the preeminent medical transport
                                    service in the Pacific Northwest.
                                </Measure>
                                <a href="http://airliftnw.org/" target="_blank">
                                    <Button>Get to know Airlift</Button>
                                </a>
                            </Box>
                        </Flex>
                    </Box>
                </Container>
            </Box>

            <RedBox>
                <Container px={5} py={5}>
                    <Box>
                        <Heading is="h2" fontSize={4} mb={5} color="white">
                            Team
                        </Heading>
                        <Flex flexWrap={["wrap", "wrap", "nowrap"]}>
                            <Box mb={3} mr={[0, 0, 5]}>
                                <Image
                                    backgroundSize="cover"
                                    borderRadius="50%"
                                    src={benjamin}
                                    width={1}
                                />
                                <Heading
                                    is="h5"
                                    fontSize={2}
                                    textAlign="center"
                                    mt={4}
                                    mb={2}
                                    color="white"
                                >
                                    Benjamin Leeds
                                </Heading>
                                <Span
                                    display="block"
                                    textAlign="center"
                                    color="white"
                                >
                                    Product Owner and Developer
                                </Span>
                            </Box>
                            <Box mb={3} mr={[0, 0, 5]}>
                                <Image
                                    backgroundSize="cover"
                                    borderRadius="50%"
                                    src={jessica}
                                    width={1}
                                />
                                <Heading
                                    is="h5"
                                    fontSize={2}
                                    textAlign="center"
                                    mt={4}
                                    mb={2}
                                    color="white"
                                >
                                    Jessica Basa
                                </Heading>
                                <Span
                                    display="block"
                                    textAlign="center"
                                    color="white"
                                >
                                    Developer
                                </Span>
                            </Box>
                            <Box mb={3} mr={[0, 0, 5]}>
                                <Image
                                    backgroundSize="cover"
                                    borderRadius="50%"
                                    src={tiffany}
                                    width={1}
                                />
                                <Heading
                                    is="h5"
                                    fontSize={2}
                                    textAlign="center"
                                    mt={4}
                                    mb={2}
                                    color="white"
                                >
                                    Tiffany Chen
                                </Heading>
                                <Span
                                    display="block"
                                    textAlign="center"
                                    color="white"
                                >
                                    Designer and Developer
                                </Span>
                            </Box>
                            <Box mb={3}>
                                <Image
                                    backgroundSize="cover"
                                    borderRadius="50%"
                                    src={vincent}
                                    width={1}
                                />
                                <Heading
                                    is="h5"
                                    fontSize={2}
                                    textAlign="center"
                                    mt={4}
                                    mb={2}
                                    color="white"
                                >
                                    Vincent van der Meulen
                                </Heading>
                                <Span
                                    display="block"
                                    textAlign="center"
                                    color="white"
                                >
                                    Designer and Developer
                                </Span>
                            </Box>
                        </Flex>
                    </Box>
                </Container>
            </RedBox>

            <Box bg="#F7F9FC" pt={5}>
                <Container width={1 / 2}>
                    <Image src={heli} width={1} />
                </Container>
                <Flex justifyContent="center">
                    <Box>
                        <Container p={5}>
                            <Heading
                                is="h2"
                                fontSize={4}
                                textAlign="center"
                                mb={2}
                            >
                                Ready for takeoff?
                            </Heading>
                            <Span fontWeight="normal">
                                Elevate is available to Airlift Northwest
                                employees
                            </Span>
                            <Flex justifyContent="center" mt={3}>
                                <Link to="/signin">
                                    <Button>Get Started</Button>
                                </Link>
                                <a href="https://www.uwmedicine.org/airlift-nw/contact-us">
                                    <ButtonOutline ml={2}>
                                        Contact Airlift
                                    </ButtonOutline>
                                </a>
                            </Flex>
                        </Container>
                    </Box>
                </Flex>
            </Box>
            <Footer />
        </div>
    );
};

export default HomePage;
