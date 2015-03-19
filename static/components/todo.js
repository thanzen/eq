
var React = require('react');

var ChatApp = React.createClass({

    render: function() {
        return (
          <div className="chatapp">
            nothing happen yet
             <textarea
                className="message-composer"
                name="message"
                value={this.state.text}
                onChange={this._onChange}
                onKeyDown={this._onKeyDown}
            />
          </div>
      );
    }

});
