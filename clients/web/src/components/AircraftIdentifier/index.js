import React, { Component } from "react";
import { Flex } from "grid-styled";

import Badge from "../Badge";
import Span from "../Span";
import Box from "../Box";

const allStatus = [
    {
        status: "OAM",
        statDesc: "On Mission",
        statColor: "#feb204"
    },
    {
        status: "OOS",
        statDesc: "Out of Service",
        statColor: "#e4002b"
    },
    {
        status: "RFM",
        statDesc: "Ready",
        statColor: "green"
    }
];

class AircraftIdentifier extends Component {
    getStatusColor(status) {
        return allStatus.find(stat => {
            return stat.status === status;
        }).statColor;
    }

    getStatusDesc(status) {
        return allStatus.find(stat => {
            return stat.status === status;
        }).statDesc;
    }

    render() {
        return (
            <Flex alignItems="center">
                <Span fontSize={this.props.fontSize} fontWeight="bold">
                    {this.props.callsign}
                </Span>
                <Box ml={2}>
                    <Span fontSize={this.props.fontSize}>
                        {this.props.nNum}
                    </Span>
                </Box>
                <Box ml={2}>
                    <Badge statusbg={this.getStatusColor(this.props.status)}>
                        {this.getStatusDesc(this.props.status)}
                    </Badge>
                </Box>
            </Flex>
        );
    }
}

export default AircraftIdentifier;
