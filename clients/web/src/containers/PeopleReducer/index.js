import React, { Component } from "react";
import { connect } from "react-redux";

import { fetchPeople } from "../../actions/people/actions";
import missionsReducer from "../../reducers/missionsReducer";

class PeopleProvider extends Component {
    componentDidMount() {
        this.props.fetchPeople();
    }

    render() {
        return this.props.render({ people: this.props.people });
    }
}

const mapStateToProps = ({ people }) => {
    return {
        people
    };
};

const mapDispatchToProps = {
    fetchPeople
};

export default connect(mapStateToProps, mapDispatchToProps)(PeopleProvider);
