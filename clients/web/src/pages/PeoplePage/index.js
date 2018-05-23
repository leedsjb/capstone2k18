import React, { Component } from "react";
import { connect } from "react-redux";
import { Flex } from "grid-styled";
import { Helmet } from "react-helmet";
import { Link } from "react-router-dom";

import Box from "../../components/Box";
import ButtonIcon from "../../components/ButtonIcon";
import DetailView from "../../components/DetailView";
import Divider from "../../components/Divider";
import FlexFillVH from "../../components/FlexFillVH";
import GroupsListItem from "../../components/GroupsListItem";
import Heading from "../../components/Heading";
import MasterView from "../../components/MasterView";
import MasterDetailView from "../../components/MasterDetailView";
import NavBar from "../../components/NavBar";
import PeopleListItem from "../../components/PeopleListItem";
import Span from "../../components/Span";
import ColoredAvatar from "../../components/ColoredAvatar";
import ScrollView from "../../components/ScrollView";
import Tab from "../../components/Tab";
import SearchBox from "../../components/SearchBox";
import TabBar from "../../components/TabBar";
import Text from "../../components/Text";
import TitleBar from "../../components/TitleBar";
import OutsideClickHandler from "../../components/OutsideClickHandler";

import { fetchPeople } from "../../actions/people/actions";
import { fetchPeopleDetail } from "../../actions/peopleDetail/actions";

import { fetchGroups } from "../../actions/groups/actions";
import { fetchGroupsDetail } from "../../actions/groupsDetail/actions";

import matchPath from "../../utils/matchPath";

class PeoplePage extends Component {
    constructor(props) {
        super(props);
        this.state = {
            queryPeople: "",
            isSearchingPeople: false,
            queryGroups: "",
            isSearchingGroups: false
        };
    }
    componentDidMount() {
        this.props.fetchPeople();
        this.props.fetchGroups();

        if (this.props.id) {
            this.props.fetchPeopleDetail(this.props.id);
        }

        if (this.props.groupID) {
            this.props.fetchGroupsDetail(this.props.groupID);
        }

        if (this.isGroupPeopleDetail()) {
            this.props.fetchGroupsDetail(this.getGroupID());
        }
    }

    componentWillReceiveProps(nextProps) {
        if (nextProps.id && nextProps.id !== this.props.id) {
            this.props.fetchPeopleDetail(nextProps.id);
        }

        if (nextProps.groupID && nextProps.groupID !== this.props.groupID) {
            this.props.fetchGroupsDetail(nextProps.groupID);
        }
    }

    getGroupID() {
        return new URLSearchParams(window.location.search).get("id");
    }

    isPeopleTab() {
        return matchPath(this.props.location.pathname, "/people");
    }

    isGroupDetailView() {
        return matchPath(this.props.location.pathname, "/groups/:groupID");
    }

    isGroupPeopleDetail() {
        return (
            new URLSearchParams(window.location.search).get("source") ===
            "groups"
        );
    }

    renderPeople() {
        return (
            <OutsideClickHandler
                handleClickOutside={() => {
                    if (this.state.isSearchingPeople) {
                        this.setState({
                            queryPeople: "",
                            isSearchingPeople: false
                        });
                        this.props.fetchPeople();
                    }
                }}
            >
                <Box py={3} px={4}>
                    <SearchBox
                        placeholder="Search all people"
                        handleChange={queryPeople => {
                            this.setState({ queryPeople }, () => {
                                this.props.fetchPeople(this.state.queryPeople);
                            });
                        }}
                        isSearching={this.state.isSearchingPeople}
                        query={this.state.queryPeople}
                        handleClear={() => {
                            this.setState({
                                queryPeople: "",
                                isSearchingPeople: false
                            });
                            this.props.fetchPeople();
                        }}
                        handleFocus={() => {
                            this.setState({ isSearchingPeople: true });
                        }}
                    />
                </Box>
                <ScrollView>{this.renderPeopleList()}</ScrollView>
            </OutsideClickHandler>
        );
    }

    renderPeopleList() {
        if (!this.props.people.pending && this.props.people.data.length > 0) {
            return this.props.people.data.map(person => {
                return (
                    <Link to={`/people/${person.id}`} key={person.id}>
                        <PeopleListItem
                            active={Number(this.props.id) === person.id ? 1 : 0}
                            person={person}
                        />
                    </Link>
                );
            });
        } else if (!this.props.people.pending) {
            return (
                <Box mt={4}>
                    <Heading is="h2" textAlign="center" fontSize={4}>
                        No People
                    </Heading>
                    <Text textAlign="center">Empty State Text</Text>
                </Box>
            );
        } else if (this.props.people.pending) {
            return <div>Loading...</div>;
        }
    }

    renderGroups() {
        return (
            <OutsideClickHandler
                handleClickOutside={() => {
                    if (this.state.isSearchingGroups) {
                        this.setState({
                            queryGroups: "",
                            isSearchingGroups: false
                        });
                        this.props.fetchGroups();
                    }
                }}
            >
                <Box py={3} px={4}>
                    <SearchBox
                        placeholder="Search all groups"
                        handleChange={queryGroups => {
                            this.setState({ queryGroups }, () => {
                                this.props.fetchGroups(this.state.queryGroups);
                            });
                        }}
                        isSearching={this.state.isSearchingGroups}
                        query={this.state.queryGroups}
                        handleClear={() => {
                            this.setState({
                                queryGroups: "",
                                isSearchingGroups: false
                            });
                            this.props.fetchGroups();
                        }}
                        handleFocus={() => {
                            this.setState({ isSearchingGroups: true });
                        }}
                    />
                </Box>
                <ScrollView>{this.renderGroupList()}</ScrollView>
            </OutsideClickHandler>
        );
    }

    renderGroupList() {
        if (!this.props.groups.pending && this.props.groups.data.length > 0) {
            return this.props.groups.data.map(group => {
                return (
                    <Link to={`/groups/${group.id}`} key={group.id}>
                        <GroupsListItem
                            active={
                                Number(this.props.groupID) === group.id ? 1 : 0
                            }
                            group={group}
                        />
                    </Link>
                );
            });
        } else if (!this.props.groups.pending) {
            return (
                <Box mt={4}>
                    <Heading is="h2" textAlign="center" fontSize={4}>
                        No Groups
                    </Heading>
                    <Text textAlign="center">Empty State Text</Text>
                </Box>
            );
        } else if (this.props.groups.pending) {
            return <div>Loading...</div>;
        }
    }

    renderGroupsDetail() {
        if (!this.props.groupID) {
            return (
                <DetailView>
                    <Box bg="gray" height="100%" />
                </DetailView>
            );
        } else if (
            !this.props.groupsDetail.pending &&
            !Array.isArray(this.props.groupsDetail.data)
        ) {
            return (
                <DetailView>
                    <Span fontWeight="bold" py={3} textAlign="center">
                        {this.props.groupsDetail.data.name}
                    </Span>
                    <Divider />
                    <ScrollView>
                        <Flex justifyContent="center" mt={5}>
                            <Flex
                                flexWrap="wrap"
                                justifyContent="space-between"
                            >
                                {this.props.groupsDetail.data.people.map(
                                    person => {
                                        return this.renderPeopleDetail(person);
                                    }
                                )}
                            </Flex>
                        </Flex>
                    </ScrollView>
                </DetailView>
            );
        } else if (!this.props.groupsDetail.pending) {
            return (
                <Box mt={4}>
                    <Heading is="h2" textAlign="center" fontSize={4}>
                        No Group Details
                    </Heading>
                    <Text textAlign="center">Empty State Text</Text>
                </Box>
            );
        } else if (this.props.groupsDetail.pending) {
            return <div>Loading...</div>;
        }
    }

    renderPeopleDetail(person) {
        if (!this.props.id && !this.props.groupID) {
            return <Box bg="gray" height="100%" />;
        } else if (
            (!this.props.peopleDetail.pending &&
                !Array.isArray(this.props.peopleDetail.data)) ||
            this.isGroupDetailView()
        ) {
            let mx = this.isGroupDetailView() ? 20 : 0;
            let mb = this.isGroupDetailView() ? 8 : 0;
            let flex = this.isGroupDetailView()
                ? ["0 1 100%", "0 1 100%", "0 1 100%", "0 1 100%", "0 1 33%"]
                : "0 1 auto";
            return (
                <Flex flex={flex} justifyContent="center" key={person.id}>
                    <Flex
                        alignItems="center"
                        flexDirection="column"
                        justifyContent="center"
                        mb={mb}
                        mx={mx}
                    >
                        <Link to={`/people/${person.id}`}>
                            <Flex
                                alignItems="center"
                                flexDirection="column"
                                justifyContent="center"
                            >
                                <Box mt={4}>
                                    <ColoredAvatar
                                        fName={person.fName}
                                        size={72}
                                    />
                                </Box>
                                <Span
                                    children={`${person.fName} ${person.lName}`}
                                    fontSize={4}
                                    fontWeight="bold"
                                    mt={3}
                                    textAlign="center"
                                />
                                <Span
                                    children={`${person.position}`}
                                    fontWeight="normal"
                                    fontSize={2}
                                />
                            </Flex>
                            <Flex mt={3}>
                                <ButtonIcon glyph="bubbleChat">Text</ButtonIcon>
                                <Box mx={3}>
                                    <ButtonIcon glyph="phone">Call</ButtonIcon>
                                </Box>
                                <ButtonIcon glyph="email">Mail</ButtonIcon>
                            </Flex>
                        </Link>
                        <Heading>Placeholder</Heading>
                    </Flex>
                </Flex>
            );
        } else {
            return null;
        }
    }

    renderMasterView() {
        if (this.props.people.error || this.props.groups.error) {
            return (
                <MasterView>
                    An error has occurred: {this.props.people.error.toString()}
                </MasterView>
            );
        } else {
            let list = this.isPeopleTab()
                ? this.renderPeople()
                : this.renderGroups();

            return (
                <MasterView>
                    <Flex>
                        <Tab
                            active={this.isPeopleTab() ? 1 : 0}
                            is={Link}
                            to="/people"
                        >
                            People
                        </Tab>
                        <Tab
                            active={!this.isPeopleTab() ? 1 : 0}
                            is={Link}
                            to="/groups"
                        >
                            Groups
                        </Tab>
                    </Flex>
                    <Divider />
                    {list}
                </MasterView>
            );
        }
    }

    renderDetailView() {
        return this.isPeopleTab() ? (
            <DetailView>
                {this.renderPeopleDetail(this.props.peopleDetail.data)}
            </DetailView>
        ) : (
            this.renderGroupsDetail()
        );
    }

    render() {
        let title = this.isPeopleTab() ? "People" : "Groups";
        return (
            <FlexFillVH flexDirection="column">
                <Helmet>
                    <title>Missions</title>
                </Helmet>

                <TitleBar title={title} />
                <NavBar />

                <MasterDetailView>
                    {this.renderMasterView()}
                    {this.renderDetailView()}
                </MasterDetailView>
                <TabBar />
            </FlexFillVH>
        );
    }
}

function mapStateToProps(state, ownProps) {
    return {
        people: state.people,
        peopleDetail: state.peopleDetail,
        groups: state.groups,
        groupsDetail: state.groupsDetail,
        location: state.router.location,
        id: ownProps.match.params.id,
        groupID: ownProps.match.params.groupID
    };
}

const mapDispatchToProps = {
    fetchPeople,
    fetchPeopleDetail,
    fetchGroups,
    fetchGroupsDetail
};

export default connect(mapStateToProps, mapDispatchToProps)(PeoplePage);
