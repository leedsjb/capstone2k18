import React from "react";
import { Link } from "react-router-dom";

import Relative from "../Relative";
import Absolute from "../Absolute";
import Card from "../Card";
import Box from "../Box";
import Span from "../Span";
import ColoredBox from "../ColoredBox";

const GroupCard = ({ groupName, to, ...props }) => {
    return (
        <Card {...props}>
            <Link to={to}>
                <Relative pt="calc(100% / 3)">
                    <Absolute top="0" right="0" bottom="0" left="0">
                        <ColoredBox word={groupName} w={1} height="100%" />
                    </Absolute>
                </Relative>
                <Box px={4} py={3}>
                    <Span fontWeight="bold">{groupName}</Span>
                </Box>
            </Link>
        </Card>
    );
};

export default GroupCard;
