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
import NavBar from "../../components/NavBar";
import TabBar from "../../components/TabBar";
import TitleBar from "../../components/TitleBar";
import Text from "../../components/Text";
import ScrollView from "../../components/ScrollView";
import SearchBox from "../../components/SearchBox";
import Span from "../../components/Span";
import OutsideClickHandler from "../../components/OutsideClickHandler";

import { fetchAircraft } from "../../actions/aircraft/actions";
import { fetchAircraftDetail } from "../../actions/aircraftDetail/actions";
import openSocket from "../../actions/socket/openSocket";

const statusFilters = ["Any status", "On Mission", "OOS"];

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
                    {aircraft.data.map(a => {
                        return (
                            <Clickable
                                key={a.id}
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
                                        Number(this.props.id) === a.id ? 1 : 0
                                    }
                                    aircraft={a}
                                />
                            </Clickable>
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
            return <div>Loading...</div>;
        }
    }

    renderAircraftDetail(aircraftDetail) {
        if (!aircraftDetail.pending) {
            return (
                <AircraftDetailListItem
                    aircraftDetail={aircraftDetail}
                    key={aircraftDetail.data.id}
                />
            );
        }
        return <div>Loading...</div>;
    }

    renderMasterView = () => {
        return this.props.aircraft.error ? (
            <div>
                An error has occurred: {this.props.aircraft.error.toString()}
            </div>
        ) : (
            <OutsideClickHandler
                handleClickOutside={() => {
                    if (this.state.isSearching) {
                        this.setState({ query: "", isSearching: false });
                        this.props.fetchAircraft();
                    }
                }}
            >
                <Box bg="#F7F8FC" px={3} py={3}>
                    <SearchBox
                        handleChange={query => {
                            this.setState({ query }, () => {
                                this.props.fetchAircraft(
                                    this.state.query,
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
                                    if (status === "Any status") {
                                        status = null;
                                    }
                                    this.props.fetchAircraft(null, status);
                                }}
                            />
                            <Box ml={3}>
                                <DropdownSelect
                                    items={statusFilters}
                                    onChange={status => {
                                        if (status === "Any status") {
                                            status = null;
                                        }
                                        this.props.fetchAircraft(null, status);
                                    }}
                                />
                            </Box>
                        </Flex>
                    ) : null}
                </Box>

                <Divider />
                <ScrollView>
                    {this.renderAircraft(this.props.aircraft)}
                </ScrollView>
            </OutsideClickHandler>
        );
    };

    renderDetailView = () => {
        return this.props.aircraftDetail.error ? (
            <div>
                An error has occurred:{" "}
                {this.props.aircraftDetail.error.toString()}
            </div>
        ) : (
            this.renderAircraftDetail(this.props.aircraftDetail)
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
