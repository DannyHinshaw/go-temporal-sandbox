# temporal

Package to hold common constants, helpers, utilities (etc.) for services to use when interacting with Temporal service.

## Registry

This package contains a "registry" for internal usage. The registry is simply the constant values of the worker services
and workflow names for usage in various other services/places. It gives central definitions that we can
change/edit/correct without having to track down string literals scattered about or, in the case of workflows, having 
to import the actual workflow to execute it, since we only really need the name of the function.

## Naming
The convention for naming Temporal task queues and namespaces is to use the name of the worker service that will be
operating on that queue/namespace. This will allow us to easily drill down into the workers running workflows in the
Temporal Web UI via namespaces, simplify task queue naming, and allow worker scaling by service (via Temporal global
namespaces), should it ever be required. This is not the only way to utilize namespaces, but it is rather a simple
example for the sake of convention. However, if you were in a multi-tenant environment, you may use namespaces to
delineate the different clients you service (for example).
