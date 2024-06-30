Hello there!

1. Namespace:
In this cool guy will be live our database!

    1.1. StatefulSet
    We have StateFulSet which describe our settings for deployment

    <Describe>
        StatefulSet is an ibnuild component of K8S. This can be used in place Deployments for declaring and executing K8S pods
        | StatefulSet manages the deployment and scaling of a set of Pods and provides guarantees about the ordering and uniquenes of these Pods.
        Unlike deployment, StatefulSet manages each pod separately ny creating separate PVC, PV and Storage classes for each pod. That is the pods use separate physical data storage.
    </Describe>


        1.1.1. Service
        This guy help us to knocn-in to our database

        1.1.2. Secret
        In this stuff we will be SCP our secrets XD

        1.1.3. ConfigMap
        Maybe in close future I will use this thing, but dont now

2. Links:
    2.1. https://chetak.hashnode.dev/database-on-kubernetes