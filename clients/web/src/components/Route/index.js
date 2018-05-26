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
                    <Relative>
                        <Flex key={waypoint.id} mt={i !== 0 ? 2 : 0}>
                            <Box
                                w={7}
                                height={7}
                                borderRadius="50%"
                                bg="black"
                                mt={2}
                                mr={2}
                            />

                            {i !== waypoints.length - 1 ? (
                                <Absolute
                                    bg="black"
                                    left={3}
                                    bottom={-16}
                                    top={8}
                                    w="1px"
                                />
                            ) : null}

                            <Flex justifyContent="space-between" flex={1}>
                                <Box w={0.6 / 1}>
                                    <Span>{waypoint.name}</Span>
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
