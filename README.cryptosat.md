
# ISS Experiment

The are two scripts in the `scripts` directory. the `drand_earth.sh` is meant
to be run on the earth node and `drand_iss.sh` is meant to be run in space. the
earth node is the leader node, which means the script for the earth node must
run **before** the script for the iss node.

The scripts initialize the nodes using a shared secret and then proceed to poll
the drand group every second to obtain random values for a total of 10 minutes.
The results are stored in logs. Note that each round is configured to be 10
seconds, therefore we over sample the group.

## Ports

The hardcoded ports in the scripts are `8654` for earth node and `8655` for the
node on the ISS. Both ports need to be open for receiving connections.

## Filesystem

The scripts assumes the directory `/drand/` exists with write access to the
`drand` process.

## Output

Outputs for all sub commands are stored in the `/drand/output` directory. Most
importantly, the file `/drand/output/iss_get.log` will contain the veriable
random values on the iss node, and the `/drand/output/earth_get.log` will
contain the same random values on the earth node.

## Execution

To execute the experiment run the following in a shell in the earth node:

    scripts/drand_earth.sh <Earth node hostname>

Then open a shell in the iss node (ordering is important) and run:

    scripts/drand_iss.sh <ISS node hostname> <Earth node hostname>

Hostname can be either a domain name or an IP address.

## Testing

All filenames and ports were chosen to prevent clashes between the earth node
and the iss node, allowing them both to run on the same machine with a single
caveat: the scripts contain an instruction to kill all prerunning drand
processes so that when the scripts run they can ensure they start out with a
clean slate. This instruction should be commented out when testing both nodes
on the same machine.

To test locally open two terminal windows. In the first terminal window run:

    scripts/drand_earth.sh localhost

In the second terminal window run:

    scripts/drand_iss.sh localhost localhost

Remeber to comment out the instruction to kill running drand processes first,
and to run the script for the earth node prior to the script for the iss node.


