import React from "react";
import { Flex } from "grid-styled";

import Box from "../Box";
import Span from "../Span";
import Icon from "../Icon";

const OOSInformation = ({ OOS }) => {
    return (
        <div>
            {OOS.reason ? (
                <Box>
                    <Span fontWeight="bold">OOS Reason: </Span>
                    <Span>{OOS.reason}</Span>
                </Box>
            ) : null}
            {OOS.remaining ? (
                <Flex mt={4} alignItems="center">
                    <Flex mr={1} alignItems="center">
                        <Icon glyph="hourglass" size={20} color="black2" />
                    </Flex>
                    <Span>{`Available in ${OOS.remaining}`} </Span>
                </Flex>
            ) : null}
            {OOS.duration ? (
                <Flex mt={4} alignItems="center">
                    <Flex mr={1} alignItems="center">
                        <Icon glyph="stopwatch" size={20} color="black2" />
                    </Flex>
                    <Span>{`${OOS.duration} elapsed`} </Span>
                </Flex>
            ) : null}
        </div>
    );
};

export default OOSInformation;
