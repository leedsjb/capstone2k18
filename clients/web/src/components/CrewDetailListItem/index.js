import React from "react";
import { Flex } from "grid-styled";
import { Link } from "react-router-dom";

import Box from "../Box";
import Card from "../Card";
import Span from "../Span";
import ColoredAvatar from "../ColoredAvatar";

const CrewDetailListItem = ({ crew, ...props }) => {
    return (
        <Flex flexWrap="wrap" justifyContent="space-between">
            {crew.map(c => {
                return (
                    <Box key={c.id} mt={4} w="calc(50% - 8px)">
                        <Link to={`/people/${c.id}`}>
                            <Card {...props} height="100%" p={5}>
                                <Flex
                                    flexDirection="column"
                                    alignItems="center"
                                    flex={1}
                                >
                                    <ColoredAvatar fName={c.fName} size={48} />
                                    <Flex
                                        mt={3}
                                        alignItems="center"
                                        flexDirection="column"
                                    >
                                        <Flex mb={1} justifyContent="center">
                                            <Span
                                                fontWeight="bold"
                                                textAlign="center"
                                            >{`${c.fName} ${c.lName}`}</Span>
                                        </Flex>
                                        <Span textAlign="center">
                                            {c.position}
                                        </Span>
                                    </Flex>
                                </Flex>
                            </Card>
                        </Link>
                    </Box>
                );
            })}
        </Flex>
    );
};

export default CrewDetailListItem;
