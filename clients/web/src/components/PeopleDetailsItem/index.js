import React from "react";
import { Link } from "react-router-dom";
import { Flex } from "grid-styled";

import Box from "../../components/Box";
import ButtonIcon from "../../components/ButtonIcon";
import Card from "../../components/Card";
import ColoredAvatar from "../../components/ColoredAvatar";
import ColoredBox from "../../components/ColoredBox";
import Container from "../../components/Container";
import Heading from "../../components/Heading";
import ProfileSnippet from "../../components/ProfileSnippet";
import Span from "../../components/Span";

const PeopleDetailsItem = ({ person }) => {
    return (
        <Box py={12}>
            <Container px={8}>
                <Flex flexDirection="column" alignItems="center">
                    <Box>
                        <ColoredAvatar fName={person.fName} size={72} />
                    </Box>
                    <Heading
                        children={`${person.fName} ${person.lName}`}
                        is="h1"
                        fontSize={4}
                        fontWeight="bold"
                        textAlign="center"
                        mt={6}
                    />
                    <Box mt={2}>
                        <Span
                            children={person.position}
                            fontWeight="normal"
                            fontSize={3}
                        />
                    </Box>
                </Flex>
                <Flex mt={6} justifyContent="center">
                    <Box>
                        <a href={`sms:${person.mobile}`}>
                            <ButtonIcon glyph="bubbleChat">Text</ButtonIcon>
                        </a>
                    </Box>
                    <Box mx={3}>
                        <a href={`tel:${person.mobile}`}>
                            <ButtonIcon glyph="phone">Call</ButtonIcon>
                        </a>
                    </Box>
                    <Box>
                        <a href={`mailto:${person.email}`}>
                            <ButtonIcon glyph="email">Mail</ButtonIcon>
                        </a>
                    </Box>
                </Flex>
                <ProfileSnippet label="Email" value={person.email} mt={12} />
                <ProfileSnippet
                    label="First name"
                    value={person.fName}
                    mt={6}
                />
                <ProfileSnippet label="Last name" value={person.lName} mt={6} />
                <ProfileSnippet
                    label="Phone"
                    value={
                        person.mobile.length === 10
                            ? `(${person.mobile.substring(
                                  0,
                                  3
                              )}) ${person.mobile.substring(
                                  3,
                                  6
                              )}-${person.mobile.substring(6, 10)}`
                            : person.mobile
                    }
                    mt={6}
                />
            </Container>
            {person.memberGroups ? (
                <div>
                    <Container px={8}>
                        <Heading fontSize={4} mt={6}>
                            Groups
                        </Heading>
                    </Container>
                    <Box maxWidth={1024} px={4} mx="auto">
                        <Flex flexWrap="wrap">
                            {person.memberGroups.map((group, i) => {
                                return (
                                    <Card
                                        key={group.id}
                                        mt={4}
                                        mx={4}
                                        w={[
                                            "calc(100% - 32px)",
                                            "calc(100% - 32px)",
                                            "calc(100% - 32px)",
                                            "calc(100% / 2 - 32px)",
                                            "calc(100% / 3 - 32px)"
                                        ]}
                                    >
                                        <Link to={`/groups/${group.id}`}>
                                            <ColoredBox
                                                word={group.name}
                                                w={1}
                                                height={64}
                                            />
                                            <Box px={4} py={3}>
                                                <Span fontWeight="bold">
                                                    {group.name}
                                                </Span>
                                            </Box>
                                        </Link>
                                    </Card>
                                );
                            })}
                        </Flex>
                    </Box>
                </div>
            ) : null}
        </Box>
    );
};

export default PeopleDetailsItem;
