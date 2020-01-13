var UserProfile = (function() {
  
    var getName = function() {
      return localStorage.getItem("user");
    };
  
    var setName = function(name) {
      localStorage.setItem("name", name);     
    };

    var getActions = function() {
      return localStorage.getItem("actions");
    };
    
    var setActions = function(actions) {
      localStorage.setItem("actions", actions);      
    };

    var includeAction = function(action) {
      return localStorage.getItem("actions").includes(action);
    };

    var setLogged = function(log) {
      localStorage.removeItem("isLogged");
      localStorage.setItem("isLogged", log);
    }

    var isLogged = function() {
        return localStorage.getItem("isLogged");
    };
  
    return {
      getName: getName,
      setName: setName,
      getActions: getActions,
      setActions: setActions,
      includeAction: includeAction,
      setLogged: setLogged,
      isLogged: isLogged
    }
  
  })();
  
  export default UserProfile;