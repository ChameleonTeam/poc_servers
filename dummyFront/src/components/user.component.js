var UserProfile = (function() {
  
    var getName = function() {
      return sessionStorage.getItem("user");
    };
  
    var setName = function(name) {
      sessionStorage.setItem("name", name);     
    };

    var getActions = function() {
      return sessionStorage.getItem("actions");
    };
    
    var setActions = function(actions) {
      sessionStorage.setItem("actions", actions);      
    };

    var includeAction = function(action) {
      return sessionStorage.getItem("actions").includes(action);
    };

    var setLogged = function(log) {
      sessionStorage.removeItem("isLogged");
      sessionStorage.setItem("isLogged", log);
    }

    var isLogged = function() {
        return sessionStorage.getItem("isLogged");
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