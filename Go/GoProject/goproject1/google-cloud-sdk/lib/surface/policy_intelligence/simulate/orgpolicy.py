# -*- coding: utf-8 -*- #
# Copyright 2022 Google LLC. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
"""Command to simulate orgpolicy."""

from __future__ import absolute_import
from __future__ import division
from __future__ import unicode_literals

from apitools.base.py import list_pager
from googlecloudsdk.api_lib.orgpolicy import utils as orgpolicy_utils
from googlecloudsdk.api_lib.policy_intelligence import orgpolicy_simulator
from googlecloudsdk.calliope import arg_parsers
from googlecloudsdk.calliope import base
from googlecloudsdk.calliope import exceptions
from googlecloudsdk.command_lib.policy_intelligence.simulator.orgpolicy import utils


_DETAILED_HELP_ALPHA = {
    'brief':
        """\
      Preview of Violations Service for OrgPolicy Simulator.
        """,
    'DESCRIPTION':
        """\
      Preview of Violations Service for OrgPolicy Simulator.
        """,
    'EXAMPLES':
        """\
      Simulate changes to Organization Policies:, run:

        $ {command}
        --policy policy.json
        --custom-constraint custom-constraint.json

      See https://cloud.google.com/iam for more information about Org Policy Simulator.
      The official Org Policy Simulator document will be released soon.

      """
}


def _ArgsAlpha(parser):
  """Parses arguments for the commands."""
  parser.add_argument(
      '--policies',
      type=arg_parsers.ArgList(),
      metavar='POLICIES',
      action=arg_parsers.UpdateAction,
      help="""Path to the JSON or YAML file that contains the Org Policy to simulate.
      Multiple Policies can be simulated by providing multiple, comma-separated paths.
      E.g. --policies=p1.json,p2.json.
      The format of policy can be found and created by `gcloud org-policies set-policy`.
      See https://cloud.google.com/sdk/gcloud/reference/org-policies/set-policy for more details.
      """)

  parser.add_argument(
      '--custom-constraints',
      type=arg_parsers.ArgList(),
      metavar='CUSTOM-CONSTRAINTS',
      action=arg_parsers.UpdateAction,
      help="""Path to the JSON or YAML file that contains the Custom Constraints to simulate.
      Multiple Custom Constraints can be simulated by providing multiple, comma-separated paths.
      e.g., --custom-constraints=constraint1.json,constraint2.json.
      """)


@base.ReleaseTracks(base.ReleaseTrack.ALPHA)
@base.Hidden
class SimulateAlpha(base.Command):
  """Simulate the Org Policies."""

  detailed_help = _DETAILED_HELP_ALPHA

  @staticmethod
  def Args(parser):
    """Parses arguments for the commands."""
    _ArgsAlpha(parser)

  def Run(self, args):
    if not args.policies and not args.custom_constraints:
      raise exceptions.ConflictingArgumentsException(
          'Must specify either --policies or --custom-constraints or both.')

    orgpolicy_simulator_api = orgpolicy_simulator.OrgPolicySimulatorApi(
        self.ReleaseTrack())

    policies = []
    for policy in args.policies:
      policy_message = utils.GetPolicyMessageFromFile(policy,
                                                      self.ReleaseTrack())
      if not policy_message.policy.name:
        raise exceptions.InvalidArgumentException(
            'Policy name',
            "'name' field not present in the organization policy.")
      policies.append(policy_message)
    custom_constraints = []
    for custom_constraint in args.custom_constraints:
      constraint_message = utils.GetCustomConstraintMessageFromFile(
          custom_constraint,
          self.ReleaseTrack())
      if not constraint_message.customConstraint.name:
        raise exceptions.InvalidArgumentException(
            'Custom constraint name',
            "'name' field not present in the custom constraint.")
      custom_constraints.append(constraint_message)

    overlay = orgpolicy_simulator_api.GetOrgPolicyOverlay(
        policies=policies, custom_constraints=custom_constraints)

    # Generate Violations Preview and get long operation id
    organization_resource = orgpolicy_utils.GetResourceFromPolicyName(
        policies[0].policy.name)
    parent = utils.GetParentFromOrganization(organization_resource)
    violations = orgpolicy_simulator_api.GetPolicysimulatorOrgPolicyViolationsPreview(
        overlay=overlay)
    request = orgpolicy_simulator_api.GenerateOrgPolicyViolationsPreviewRequest(
        violations_preview=violations,
        parent=parent)
    op_simulator_service = orgpolicy_simulator_api.client.OrganizationsLocationsService(
        orgpolicy_simulator_api.client)
    violations_preview_operation = op_simulator_service.OrgPolicyViolationsPreviews(
        request=request)

    # Poll Long Running Operation and get Violations Name
    _ = orgpolicy_simulator_api.WaitForOperation(
        violations_preview_operation,
        'Waiting for operation [{}] to complete'.format(
            violations_preview_operation.name))

    # List results of the violations_preview.
    list_violations_preview_request = orgpolicy_simulator_api.messages.PolicysimulatorOrganizationsLocationsOrgPolicyViolationsPreviewsListRequest(
        parent=parent)
    pov_service = orgpolicy_simulator_api.client.OrganizationsLocationsOrgPolicyViolationsPreviewsService(
        orgpolicy_simulator_api.client)

    return list_pager.YieldFromList(
        pov_service,
        list_violations_preview_request,
        batch_size=1000,
        field='orgPolicyViolationsPreviews',
        batch_size_attribute='pageSize')
