/* Imports: External */
import { DeployFunction } from 'hardhat-deploy/dist/types'
import { hexStringEquals, awaitCondition } from '@mantleio/core-utils'

/* Imports: Internal */
import {
  deployAndVerifyAndThen,
  getContractFromArtifact,
} from '../src/deploy-utils'
import { names } from '../src/address-names'

const deployFn: DeployFunction = async (hre) => {
  const Lib_AddressManager = await getContractFromArtifact(
    hre,
    names.unmanaged.Lib_AddressManager
  )

  await deployAndVerifyAndThen({
    hre,
    name: names.managed.contracts.BVM_L1CrossDomainMessenger,
    contract: 'L1CrossDomainMessenger',
    args: [],
    postDeployAction: async (contract) => {
      // Theoretically it's not necessary to initialize this contract since it sits behind
      // a proxy. However, it's best practice to initialize it anyway just in case there's
      // some unknown security hole. It also prevents another user from appearing like an
      // official address because it managed to call the initialization function.
      console.log(`Initializing L1CrossDomainMessenger (implementation)...`)
      await contract.initialize(Lib_AddressManager.address)

      console.log(`Checking that contract was correctly initialized...`)
      await awaitCondition(
        async () => {
          return hexStringEquals(
            await contract.libAddressManager(),
            Lib_AddressManager.address
          )
        },
        5000,
        100
      )

      // Same thing as above, we want to transfer ownership of this contract to the owner of the
      // AddressManager. Not technically necessary but seems like the right thing to do.
      console.log(
        `set pauseowner of L1CrossDomainMessenger (implementation)...`
      )
      // set the pauseowner
      const pauseowner = hre.deployConfig.bvmCrossDomainPauseOwner
      await contract.setPauseOwner(pauseowner)
      console.log(`Checking that contract pause owner was correctly set...`)
      await awaitCondition(
        async () => {
          console.log(contract.getPauseOwner())
          console.log(pauseowner)
          return hexStringEquals(await contract.getPauseOwner(), pauseowner)
        },
        5000,
        100
      )

      // Same thing as above, we want to transfer ownership of this contract to the owner of the
      // AddressManager. Not technically necessary but seems like the right thing to do.
      console.log(
        `Transferring ownership of L1CrossDomainMessenger (implementation)...`
      )
      const owner = hre.deployConfig.bvmAddressManagerOwner
      await contract.transferOwnership(owner)

      console.log(`Checking that contract owner was correctly set...`)
      await awaitCondition(
        async () => {
          return hexStringEquals(await contract.owner(), owner)
        },
        5000,
        100
      )
    },
  })
}

deployFn.tags = ['L1CrossDomainMessenger', 'upgrade']

export default deployFn
