import React, { Component } from "react";
import { connect } from "react-redux";
import { Flex } from "grid-styled";
import { Helmet } from "react-helmet";
import { push } from "react-router-redux";

import AircraftListItem from "../../components/AircraftListItem";
import AircraftDetailListItem from "../../components/AircraftDetailListItem";
import Box from "../../components/Box";
import Clickable from "../../components/Clickable";
import Divider from "../../components/Divider";
import DropdownSelect from "../../components/DropdownSelect";
import FlexFillVH from "../../components/FlexFillVH";
import Heading from "../../components/Heading";
import MasterDetailMapView from "../../components/MasterDetailMapView";
import AircraftLoader from "../../components/AircraftLoader";
import NavBar from "../../components/NavBar";
import TabBar from "../../components/TabBar";
import TitleBar from "../../components/TitleBar";
import Text from "../../components/Text";
import ScrollView from "../../components/ScrollView";
import SearchBox from "../../components/SearchBox";
import OutsideClickHandler from "../../components/OutsideClickHandler";
import LoadingSpinner from "../../components/LoadingSpinner";

import { fetchAircraft } from "../../actions/aircraft/actions";
import { fetchAircraftDetail } from "../../actions/aircraftDetail/actions";
import openSocket from "../../actions/socket/openSocket";

const AS = "Any status";
const OAM = "On a mission";
const RFM = "Ready for mission";
const OOS = "OOS";

const AC = "Any category";
const FW = "Fixed-wing";
const RC = "Rotorcraft";

const statusFilters = [AS, RFM, OAM, OOS];
const categoryFilters = [AC, FW, RC];

class AircraftPage extends Component {
    constructor(props) {
        super(props);
        this.state = {
            query: "",
            isSearching: false
        };
    }

    componentDidMount() {
        this.props.fetchAircraft();
        if (this.props.id) {
            this.props.fetchAircraftDetail(this.props.id);
        }
        // this.props.openSocket();
    }

    componentWillReceiveProps(nextProps) {
        if (nextProps.id && nextProps.id !== this.props.id) {
            this.props.fetchAircraftDetail(nextProps.id);
        }
    }

    renderAircraft(aircraft) {
        if (!aircraft.pending && aircraft.data.length > 0) {
            return (
                <div>
                    {aircraft.data.map((a, i) => {
                        return (
                            <div key={a.id}>
                                <Clickable
                                    onClick={() => {
                                        if (this.state.isSearching) {
                                            if (this.state.isSearching) {
                                                this.setState({
                                                    query: "",
                                                    isSearching: false
                                                });
                                                this.props.fetchAircraft();
                                            }
                                        }
                                        this.props.push(`/aircraft/${a.id}`);
                                    }}
                                >
                                    <AircraftListItem
                                        active={
                                            Number(this.props.id) === a.id
                                                ? 1
                                                : 0
                                        }
                                        aircraft={a}
                                    />
                                </Clickable>
                                {aircraft.data.length === 1 ||
                                i !== aircraft.data.length - 1 ? (
                                    <Divider />
                                ) : null}
                            </div>
                        );
                    })}
                </div>
            );
        } else if (!aircraft.pending) {
            return (
                <Box mt={4}>
                    <Heading is="h2" textAlign="center" fontSize={4}>
                        No Aircraft
                    </Heading>
                    <Text textAlign="center">Empty State Text</Text>
                </Box>
            );
        } else {
            return (
                <div>
                    <AircraftLoader />
                    <AircraftLoader />
                    <AircraftLoader />
                </div>
            );
        }
    }

    renderAircraftDetail(aircraftDetail) {
        if (!aircraftDetail.pending && !Array.isArray(aircraftDetail.data)) {
            return (
                <AircraftDetailListItem
                    aircraftDetail={aircraftDetail}
                    key={aircraftDetail.data.id}
                />
            );
        }
        return <LoadingSpinner />;
    }

    renderMasterView = () => {
        return this.props.aircraft.error ? (
            <Flex
                flexDirection="column"
                flex={1}
                justifyContent="center"
                px={4}
            >
                An error has occurred: {this.props.aircraft.error.toString()}
            </Flex>
        ) : (
            <OutsideClickHandler
                handleClickOutside={() => {
                    if (this.state.isSearching) {
                        this.setState({ query: "", isSearching: false });
                        this.props.fetchAircraft();
                    }
                }}
            >
                <Box
                    px={3}
                    py={3}
                    boxShadow="0 2px 8px #EBEBEB"
                    borderBottom="1px solid #E0E0E0"
                    position="relative"
                    zIndex={999}
                >
                    <SearchBox
                        handleChange={query => {
                            this.setState({ query }, () => {
                                this.props.fetchAircraft(
                                    this.state.query,
                                    null,
                                    null
                                );
                            });
                        }}
                        isSearching={this.state.isSearching}
                        query={this.state.query}
                        handleClear={() => {
                            this.setState({ query: "", isSearching: false });
                            this.props.fetchAircraft();
                        }}
                        placeholder="Search all aircraft"
                        handleFocus={() => {
                            this.setState({ isSearching: true });
                        }}
                    />

                    {!this.state.isSearching ? (
                        <Flex alignItems="center" mt={2}>
                            <DropdownSelect
                                items={statusFilters}
                                onChange={status => {
                                    switch (status) {
                                        case OAM:
                                            status = "oam";
                                            break;
                                        case RFM:
                                            status = "rfm";
                                            break;
                                        case OOS:
                                            status = "oos";
                                            break;
                                        default:
                                            status = null;
                                            break;
                                    }

                                    this.props.fetchAircraft(
                                        null,
                                        status,
                                        null
                                    );
                                }}
                            />
                            <Box ml={3}>
                                <DropdownSelect
                                    items={categoryFilters}
                                    onChange={category => {
                                        switch (category) {
                                            case FW:
                                                category = "fixed-wing";
                                                break;
                                            case RC:
                                                category = "rotorcraft";
                                                break;
                                            default:
                                                category = "";
                                                break;
                                        }

                                        this.props.fetchAircraft(
                                            null,
                                            null,
                                            category
                                        );
                                    }}
                                />
                            </Box>
                        </Flex>
                    ) : null}
                </Box>

                <ScrollView>
                    {this.renderAircraft(this.props.aircraft)}
                </ScrollView>
            </OutsideClickHandler>
        );
    };

    renderDetailView = () => {
        return (
            <Flex flexDirection="column" flex={1}>
                {this.props.aircraftDetail.error ? (
                    <div>
                        An error has occurred:{" "}
                        {this.props.aircraftDetail.error.toString()}
                    </div>
                ) : (
                    this.renderAircraftDetail(this.props.aircraftDetail)
                )}
            </Flex>
        );
    };

    render() {
        return (
            <FlexFillVH flexDirection="column">
                <Helmet>
                    <title>Aircraft</title>
                </Helmet>
                <TitleBar title="Aircraft" showMap link="/aircraft/map" />
                <NavBar />
                <MasterDetailMapView
                    renderMasterView={this.renderMasterView}
                    renderDetailView={this.renderDetailView}
                    showDetail={this.props.id}
                />
                <TabBar />
            </FlexFillVH>
        );
    }
}

function mapStateToProps(state, ownProps) {
    return {
        aircraft: state.aircraft,
        aircraftDetail: state.aircraftDetail,
        id: ownProps.match.params.id
    };
}

const mapDispatchToProps = {
    fetchAircraft,
    fetchAircraftDetail,
    push,
    openSocket
};

export default connect(mapStateToProps, mapDispatchToProps)(AircraftPage);
