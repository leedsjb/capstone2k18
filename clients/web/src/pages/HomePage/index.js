import React from "react";
import { Flex } from "grid-styled";
import { Link } from "react-router-dom";

import Box from "../../components/Box";
import ButtonPrimaryGradient from "../../components/ButtonPrimaryGradient";
import ButtonPrimaryOutline from "../../components/ButtonPrimaryOutline";
import Container from "../../components/Container";
import Footer from "../../components/Footer";
import Heading from "../../components/Heading";
import Relative from "../../components/Relative";
import Icon from "../../components/Icon";
import Image from "../../components/Image";
import Measure from "../../components/Measure";
import RedBox from "../../components/RedBox";
import Span from "../../components/Span";

import benjamin from "../../images/benjamin.png";
import heli from "../../images/heli.png";
import logotext from "../../images/logotext.svg";
import jessica from "../../images/jessica.png";
import partner from "../../images/partner.png";
import phone from "../../images/phone.png";
import notification from "../../images/notification.svg";
import search from "../../images/search.png";
import second from "../../images/second.png";
import tiffany from "../../images/tiffany.png";
import trackMissions from "../../images/trackMissions.png";
import vincent from "../../images/vincent.png";

const HomePage = () => {
    return (
        <div>
            <Box bg="#F7F9FA">
                <Container>
                    <Box
                        mx="auto"
                        width={["100%", 488, "100%"]}
                        pt={12}
                        pb={24}
                    >
                        <Flex
                            justifyContent="space-between"
                            alignItems="center"
                        >
                            <Relative left="-3px">
                                <Link to="/">
                                    <Image
                                        src={logotext}
                                        alt="AirliftNW Elevate"
                                        w={184}
                                    />
                                </Link>
                            </Relative>
                            <div>
                                <ButtonPrimaryGradient>
                                    <Link to="/aircraft">Sign in</Link>
                                </ButtonPrimaryGradient>
                            </div>
                        </Flex>

                        <Flex
                            justifyContent="space-between"
                            alignItems="center"
                            mt={24}
                        >
                            <Flex flexDirection="column">
                                <Heading
                                    is="h1"
                                    mt={0}
                                    fontSize={6}
                                    maxWidth={720}
                                >
                                    Critical Communication for Airborne Medical
                                    Missions
                                </Heading>
                                <Measure mt={4} mb={0}>
                                    Elevate is a mission-critical,
                                    high-availability application for Airlift
                                    Northwest emergency flight nurses and pilots
                                    transporting severely ill or injured
                                    patients by helicopter to reach life-saving
                                    medical treatment.
                                </Measure>
                                <Link to="/aircraft">
                                    <ButtonPrimaryGradient mt={6}>
                                        Get Started
                                    </ButtonPrimaryGradient>
                                </Link>
                            </Flex>
                            <Box maxWidth={320}>
                                <Image src={phone} w="100%" />
                            </Box>
                        </Flex>
                    </Box>
                </Container>
            </Box>

            <RedBox>
                <Container py={24}>
                    <Flex
                        flexWrap={["wrap", "wrap", "nowrap"]}
                        alignItems="center"
                        justifyContent="center"
                    >
                        <Box maxWidth={488} mr={[0, 0, 12]}>
                            <Image src={second} width={1} />
                        </Box>
                        <Box>
                            <Heading is="h2" fontSize={5} pt={5} color="white">
                                Because every second counts
                            </Heading>
                            <Measure mt={3} color="white">
                                Elevate bridges the communication gap between
                                the dispatch center on the ground and flight
                                crews in the air by displaying patient,
                                aircraft, and aircrew status. This ensures that
                                aircraft and aircrews quickly arrive at the
                                right location with the information needed to
                                immediately administer life-saving patient care.
                            </Measure>
                        </Box>
                    </Flex>
                </Container>
            </RedBox>

            <Container py={24}>
                <Box mx="auto" width={["100%", 488, "100%"]}>
                    <Heading is="h2" fontSize={5} mb={4}>
                        Elevating health care technology
                    </Heading>
                    <Flex flexWrap={["wrap", "wrap", "nowrap"]} mt={18}>
                        <Box mr={12}>
                            <Icon glyph="devices" size={48} />
                            <Heading is="h4" fontSize={4} mt={8}>
                                Use anywhere
                            </Heading>
                            <Measure mt={3}>
                                Elevate is available on any device with a web
                                browser.
                            </Measure>
                        </Box>

                        <Box mr={12}>
                            <Icon glyph="checkShield" size={48} />
                            <Heading is="h4" fontSize={4} mt={8}>
                                HIPAA Compliant
                            </Heading>
                            <Measure mt={3}>
                                We comply with HIPAA and do not store or show
                                any PHI.
                            </Measure>
                        </Box>

                        <Box>
                            <Icon glyph="wifiCheck" size={48} />
                            <Heading is="h4" fontSize={4} mt={8}>
                                Works offline
                            </Heading>
                            <Measure mt={3}>
                                If you lose your internet connnection, the most
                                recently updated information will still be
                                available to you.
                            </Measure>
                        </Box>
                    </Flex>
                </Box>
            </Container>

            <RedBox>
                <Container py={24}>
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
                            <Heading is="h2" fontSize={5} color="white">
                                Track missions
                            </Heading>
                            <Measure mt={3} color="white">
                                Know what is going on at all times and follow
                                missions on a map as they are happening.
                            </Measure>
                        </Box>
                        <Box maxWidth={320}>
                            <Image src={trackMissions} width="100%" />
                        </Box>
                    </Flex>
                </Container>
            </RedBox>

            <Box bg="#F7F9FA">
                <Container py={24}>
                    <Box mx="auto" width={["100%", 488, "100%"]}>
                        <Flex
                            flexWrap={["wrap", "wrap", "nowrap"]}
                            alignItems="center"
                        >
                            <Box maxWidth={488} mr={[0, 0, 5]}>
                                <Image src={notification} width={1} />
                            </Box>
                            <Box>
                                <Heading is="h2" fontSize={5}>
                                    Get notified
                                </Heading>
                                <Measure mt={3}>
                                    Receive a text notification as soon as you
                                    are assigned to a new mission.
                                </Measure>
                            </Box>
                        </Flex>
                    </Box>
                </Container>
            </Box>

            <RedBox>
                <Container py={24}>
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
                            <Heading is="h2" fontSize={5} color="white">
                                Find the right person
                            </Heading>
                            <Measure mt={3} color="white">
                                Find and contact the person or group you’re
                                looking for in seconds.
                            </Measure>
                        </Box>
                        <Box maxWidth={488}>
                            <Image src={search} width={1} />
                        </Box>
                    </Flex>
                </Container>
            </RedBox>

            <Box>
                <Container py={24}>
                    <Flex
                        flexWrap={["wrap", "wrap", "nowrap"]}
                        alignItems="center"
                        justifyContent="center"
                    >
                        <Box maxWidth={488} mr={[0, 0, 12]}>
                            <Image src={partner} width={1} />
                        </Box>
                        <Box>
                            <Heading is="h2" fontSize={5}>
                                About Airlift Northwest
                            </Heading>
                            <Measure mt={3}>
                                Elevate was built together with Airlift
                                Northwest, the preeminent medical transport
                                service in the Pacific Northwest.
                            </Measure>
                            <a
                                href="http://airliftnw.org/"
                                rel="noopener noreferrer"
                                target="_blank"
                            >
                                <ButtonPrimaryGradient mt={6}>
                                    Get to know Airlift
                                </ButtonPrimaryGradient>
                            </a>
                        </Box>
                    </Flex>
                </Container>
            </Box>

            <RedBox>
                <Container py={24}>
                    <Heading is="h2" fontSize={5} color="white">
                        Team
                    </Heading>
                    <Flex
                        flexWrap="wrap"
                        justifyContent="space-between"
                        mt={18}
                    >
                        <Flex
                            flexDirection="column"
                            alignItems="center"
                            width="calc(100% / 4 - 24px)"
                        >
                            <Image
                                backgroundSize="cover"
                                borderRadius="50%"
                                src={benjamin}
                                width={160}
                            />
                            <Heading
                                fontSize={4}
                                textAlign="center"
                                mt={6}
                                mb={1}
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
                        </Flex>
                        <Flex
                            flexDirection="column"
                            alignItems="center"
                            width="calc(100% / 4 - 24px)"
                        >
                            <Image
                                backgroundSize="cover"
                                borderRadius="50%"
                                src={jessica}
                                width={160}
                            />
                            <Heading
                                is="h5"
                                fontSize={4}
                                textAlign="center"
                                mt={6}
                                mb={1}
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
                        </Flex>
                        <Flex
                            flexDirection="column"
                            alignItems="center"
                            width="calc(100% / 4 - 24px)"
                        >
                            {" "}
                            <Image
                                backgroundSize="cover"
                                borderRadius="50%"
                                src={tiffany}
                                width={160}
                            />
                            <Heading
                                is="h5"
                                fontSize={4}
                                textAlign="center"
                                mt={6}
                                mb={1}
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
                        </Flex>
                        <Flex
                            flexDirection="column"
                            alignItems="center"
                            width="calc(100% / 4 - 24px)"
                        >
                            {" "}
                            <Image
                                backgroundSize="cover"
                                borderRadius="50%"
                                src={vincent}
                                width={160}
                            />
                            <Heading
                                is="h5"
                                fontSize={4}
                                textAlign="center"
                                mt={6}
                                mb={1}
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
                        </Flex>
                    </Flex>
                </Container>
            </RedBox>

            <Box bg="#F7F9FA" py={24}>
                <Container>
                    <Flex flexDirection="column" alignItems="center">
                        <Box maxWidth={792}>
                            <Image src={heli} width={1} />
                        </Box>
                        <Heading
                            is="h2"
                            fontSize={5}
                            textAlign="center"
                            mt={4}
                            mb={2}
                        >
                            Ready for takeoff?
                        </Heading>
                        <Span fontWeight="normal" textAlign="center">
                            Elevate is available to Airlift Northwest employees
                        </Span>

                        <Flex justifyContent="center" mt={8}>
                            <Link to="/ai">
                                <ButtonPrimaryGradient>
                                    Get Started
                                </ButtonPrimaryGradient>
                            </Link>
                            <a
                                href="https://www.uwmedicine.org/airlift-nw/contact-us"
                                rel="noopener noreferrer"
                                target="_blank"
                            >
                                <ButtonPrimaryOutline ml={4}>
                                    Contact Airlift
                                </ButtonPrimaryOutline>
                            </a>
                        </Flex>
                    </Flex>
                </Container>
            </Box>
            <Footer />
        </div>
    );
};

export default HomePage;
