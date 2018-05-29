import React, { Component } from "react";

// TODO: Figure out if this is indeed the best way to
// inform children about their list index
// https://mxstbr.blog/2017/02/react-children-deepdive/
class Accordion extends Component {
    renderChildren() {
        return React.Children.map(this.props.children, (child, i) => {
            return React.cloneElement(child, {
                id: i
            });
        });
    }

    render() {
        return this.renderChildren();
    }
}

export default Accordion;
