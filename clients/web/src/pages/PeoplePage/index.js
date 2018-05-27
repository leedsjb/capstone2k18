import React, { Component } from "react";
import { connect } from "react-redux";
import { Flex } from "grid-styled";
import { Helmet } from "react-helmet";
import { Link } from "react-router-dom";

import Box from "../../components/Box";
import ButtonIcon from "../../components/ButtonIcon";
import Card from "../../components/Card";
import ColoredAvatar from "../../components/ColoredAvatar";
import Container from "../../components/Container";
import DetailView from "../../components/DetailView";
import Divider from "../../components/Divider";
import FlexFillVH from "../../components/FlexFillVH";
import GroupsListItem from "../../components/GroupsListItem";
import GroupsLoader from "../../components/GroupsLoader";
import Heading from "../../components/Heading";
import MasterView from "../../components/MasterView";
import LoadingSpinner from "../../components/LoadingSpinner";
import MasterDetailView from "../../components/MasterDetailView";
import NavBar from "../../components/NavBar";
import OutsideClickHandler from "../../components/OutsideClickHandler";
import PeopleDetailsItem from "../../components/PeopleDetailsItem";
import PeopleListItem from "../../components/PeopleListItem";
<<<<<<< HEAD
import PeopleLoader from "../../components/PeopleLoader";
import SearchBox from "../../components/SearchBox";
=======
import GroupCard from "../../components/GroupCard";
>>>>>>> development
import Span from "../../components/Span";
import ScrollView from "../../components/ScrollView";
import Tab from "../../components/Tab";
import TabBar from "../../components/TabBar";
import Text from "../../components/Text";
import TitleBar from "../../components/TitleBar";
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
            return this.props.people.data.map((person, i) => {
                return (
                    <div>
                        <Link to={`/people/${person.id}`} key={person.id}>
                            <PeopleListItem
                                active={
                                    Number(this.props.id) === person.id ? 1 : 0
                                }
                                person={person}
                            />
                        </Link>
                        {this.props.people.data.length === 1 ||
                        i !== this.props.people.data.length - 1 ? (
                            <Divider />
                        ) : null}
                    </div>
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
            return (
                <div>
                    <PeopleLoader />
                    <PeopleLoader />
                    <PeopleLoader />
                </div>
            );
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
            return this.props.groups.data.map((group, i) => {
                return (
                    <div>
                        <Link to={`/groups/${group.id}`} key={group.id}>
                            <GroupsListItem
                                active={
                                    Number(this.props.groupID) === group.id
                                        ? 1
                                        : 0
                                }
                                group={group}
                            />
                        </Link>
                        {this.props.groups.data.length === 1 ||
                        i !== this.props.groups.data.length - 1 ? (
                            <Divider />
                        ) : null}
                    </div>
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
            return (
                <div>
                    <GroupsLoader />
                    <GroupsLoader />
                    <GroupsLoader />
                </div>
            );
        }
    }

    renderGroupsDetail() {
        if (!this.props.groupID) {
            return <Box bg="gray" height="100%" />;
        } else if (
            !this.props.groupsDetail.pending &&
            !Array.isArray(this.props.groupsDetail.data)
        ) {
            return (
                <Box>
                    <Flex py={3} justifyContent="center">
                        <Span fontWeight="bold" textAlign="center">
                            {this.props.groupsDetail.data.name}
                        </Span>
                    </Flex>
                    <Divider />
                    <Container px={8} mt={6}>
                        <Heading is="h2" fontSize={4}>
                            Group members
                        </Heading>
                    </Container>
                    <Box maxWidth={1024} mx="auto" w={1} px={4}>
                        <Flex flexWrap="wrap" justifyContent="flex-start">
                            {this.props.groupsDetail.data.people.map(person => {
                                return (
                                    <Card
                                        p={4}
                                        mx={4}
                                        w={[
                                            "calc(100% / 2 - 32px)",
                                            "calc(100% - 32px)",
                                            "calc(100% - 32px)",
                                            "calc(100% / 3 - 32px)"
                                        ]}
                                        mt={4}
                                    >
                                        <Flex
                                            flexDirection="column"
                                            alignItems="center"
                                        >
                                            <ColoredAvatar
                                                fName={person.fName}
                                                size={72}
                                            />
                                            <Flex
                                                flexDirection="column"
                                                alignItems="center"
                                                mt={4}
                                            >
                                                <Span
                                                    fontWeight="bold"
                                                    textAlign="center"
                                                >
                                                    {`${person.fName} ${
                                                        person.lName
                                                    }`}
                                                </Span>
                                                <Flex
                                                    flexDirection="column"
                                                    alignItems="center"
                                                    mt={1}
                                                >
                                                    <Span textAlign="center">
                                                        {person.position}
                                                    </Span>
                                                </Flex>
                                            </Flex>
                                            <Flex
                                                flexWrap="wrap"
                                                justifyContent="center"
                                                mt={2}
                                            >
                                                <Box mt={4}>
                                                    <ButtonIcon glyph="bubbleChat">
                                                        Text
                                                    </ButtonIcon>
                                                </Box>
                                                <Box mx={3} mt={4}>
                                                    <ButtonIcon glyph="phone">
                                                        Call
                                                    </ButtonIcon>
                                                </Box>
                                                <Box mt={4}>
                                                    <ButtonIcon glyph="email">
                                                        Mail
                                                    </ButtonIcon>
                                                </Box>
                                            </Flex>
                                        </Flex>
                                    </Card>
                                );
                            })}
                        </Flex>
                    </Box>
                </Box>
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
            return <LoadingSpinner />;
        }
    }

    renderPeopleDetail() {
        if (!this.props.id && !this.props.groupID) {
            return <Box bg="gray" height="100%" />;
        } else if (
            !this.props.peopleDetail.pending &&
            !Array.isArray(this.props.peopleDetail.data)
        ) {
<<<<<<< HEAD
            return <PeopleDetailsItem person={this.props.peopleDetail.data} />;
=======
            const person = this.props.peopleDetail.data;

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
                                <ButtonIcon glyph="bubbleChat">Text</ButtonIcon>
                            </Box>
                            <Box mx={3}>
                                <ButtonIcon glyph="phone">Call</ButtonIcon>
                            </Box>
                            <Box>
                                <ButtonIcon glyph="email">Mail</ButtonIcon>
                            </Box>
                        </Flex>
                        <ProfileSnippet
                            label="Email"
                            value={person.email}
                            mt={12}
                        />
                        <ProfileSnippet
                            label="First name"
                            value={person.fName}
                            mt={6}
                        />
                        <ProfileSnippet
                            label="Last name"
                            value={person.lName}
                            mt={6}
                        />
                        <ProfileSnippet
                            label="Phone"
                            value={person.mobile}
                            mt={6}
                        />
                    </Container>
                    {person.memberGroups.length > 0 ? (
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
                                            <GroupCard
                                                mx={4}
                                                w={[
                                                    "calc(100% / 2 - 32px)",
                                                    "calc(100% - 32px)",
                                                    "calc(100% - 32px)",
                                                    "calc(100% / 3 - 32px)"
                                                ]}
                                                mt={4}
                                                groupName={group.name}
                                                to={`/groups/${group.id}`}
                                            />
                                        );
                                    })}
                                </Flex>
                            </Box>
                        </div>
                    ) : null}
                </Box>
            );
>>>>>>> development
        } else {
            return <div>Loading...</div>;
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
                    {this.isPeopleTab()
                        ? this.renderPeople()
                        : this.renderGroups()}
                </MasterView>
            );
        }
    }

    renderDetailView() {
        return (
            <DetailView>
                {this.isPeopleTab()
                    ? this.renderPeopleDetail()
                    : this.renderGroupsDetail()}
            </DetailView>
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
