import React from "react";
import { Flex } from "grid-styled";
import moment from "moment";

import Relative from "../Relative";
import Absolute from "../Absolute";
import Box from "../Box";
import Span from "../Span";

const Route = ({ waypoints }) => {
    return (
        <div>
            {waypoints.map((waypoint, i) => {
                return (
                    <Relative key={waypoint.id}>
                        <Flex mt={i !== 0 ? 2 : 0}>
                            {waypoints[i + 1] && waypoints[i + 1].completed ? (
                                <Absolute
                                    bg="black"
                                    left={3}
                                    bottom={-16}
                                    top={8}
                                    w="1px"
                                    zIndex={-1}
                                />
                            ) : null}

                            {waypoints[i + 1] && waypoints[i + 1].active ? (
                                <Absolute
                                    left={3}
                                    bottom={-16}
                                    top={8}
                                    border="0.5px dotted black"
                                    zIndex={-1}
                                />
                            ) : null}

                            <Box
                                w={7}
                                height={7}
                                borderRadius="50%"
                                bg={waypoint.completed ? "black" : "white"}
                                border="1px solid black"
                                mt={2}
                                mr={2}
                            />

                            <Flex justifyContent="space-between" flex={1}>
                                <Box w={0.6 / 1}>
                                    <Span
                                        fontWeight={
                                            waypoint.active ? "bold" : "normal"
                                        }
                                    >
                                        {waypoint.name}
                                    </Span>
                                </Box>
                                <Flex justifyContent="flex-end" w={0.3 / 1}>
                                    <Span>{moment(waypoint.ETA).toNow()}</Span>
                                </Flex>
                            </Flex>
                        </Flex>
                    </Relative>
                );
            })}
        </div>
    );
};

export default Route;
