"""Generated client library for iam version v2."""
# NOTE: This file is autogenerated and should not be edited by hand.

from __future__ import absolute_import

from apitools.base.py import base_api
from googlecloudsdk.third_party.apis.iam.v2 import iam_v2_messages as messages


class IamV2(base_api.BaseApiClient):
  """Generated client library for service iam version v2."""

  MESSAGES_MODULE = messages
  BASE_URL = 'https://iam.googleapis.com/'
  MTLS_BASE_URL = 'https://iam.mtls.googleapis.com/'

  _PACKAGE = 'iam'
  _SCOPES = ['https://www.googleapis.com/auth/cloud-platform']
  _VERSION = 'v2'
  _CLIENT_ID = 'CLIENT_ID'
  _CLIENT_SECRET = 'CLIENT_SECRET'
  _USER_AGENT = 'google-cloud-sdk'
  _CLIENT_CLASS_NAME = 'IamV2'
  _URL_VERSION = 'v2'
  _API_KEY = None

  def __init__(self, url='', credentials=None,
               get_credentials=True, http=None, model=None,
               log_request=False, log_response=False,
               credentials_args=None, default_global_params=None,
               additional_http_headers=None, response_encoding=None):
    """Create a new iam handle."""
    url = url or self.BASE_URL
    super(IamV2, self).__init__(
        url, credentials=credentials,
        get_credentials=get_credentials, http=http, model=model,
        log_request=log_request, log_response=log_response,
        credentials_args=credentials_args,
        default_global_params=default_global_params,
        additional_http_headers=additional_http_headers,
        response_encoding=response_encoding)
    self.policies_operations = self.PoliciesOperationsService(self)
    self.policies = self.PoliciesService(self)
    self.v2 = self.V2Service(self)

  class PoliciesOperationsService(base_api.BaseApiService):
    """Service class for the policies_operations resource."""

    _NAME = 'policies_operations'

    def __init__(self, client):
      super(IamV2.PoliciesOperationsService, self).__init__(client)
      self._upload_configs = {
          }

    def Get(self, request, global_params=None):
      r"""Gets the latest state of a long-running operation. Clients can use this method to poll the operation result at intervals as recommended by the API service.

      Args:
        request: (IamPoliciesOperationsGetRequest) input message
        global_params: (StandardQueryParameters, default: None) global arguments
      Returns:
        (GoogleLongrunningOperation) The response message.
      """
      config = self.GetMethodConfig('Get')
      return self._RunMethod(
          config, request, global_params=global_params)

    Get.method_config = lambda: base_api.ApiMethodInfo(
        flat_path='v2/policies/{policiesId}/{policiesId1}/{policiesId2}/operations/{operationsId}',
        http_method='GET',
        method_id='iam.policies.operations.get',
        ordered_params=['name'],
        path_params=['name'],
        query_params=[],
        relative_path='v2/{+name}',
        request_field='',
        request_type_name='IamPoliciesOperationsGetRequest',
        response_type_name='GoogleLongrunningOperation',
        supports_download=False,
    )

  class PoliciesService(base_api.BaseApiService):
    """Service class for the policies resource."""

    _NAME = 'policies'

    def __init__(self, client):
      super(IamV2.PoliciesService, self).__init__(client)
      self._upload_configs = {
          }

    def CreatePolicy(self, request, global_params=None):
      r"""Creates a policy.

      Args:
        request: (IamPoliciesCreatePolicyRequest) input message
        global_params: (StandardQueryParameters, default: None) global arguments
      Returns:
        (GoogleLongrunningOperation) The response message.
      """
      config = self.GetMethodConfig('CreatePolicy')
      return self._RunMethod(
          config, request, global_params=global_params)

    CreatePolicy.method_config = lambda: base_api.ApiMethodInfo(
        flat_path='v2/policies/{policiesId}/{policiesId1}',
        http_method='POST',
        method_id='iam.policies.createPolicy',
        ordered_params=['parent'],
        path_params=['parent'],
        query_params=['policyId'],
        relative_path='v2/{+parent}',
        request_field='googleIamV2Policy',
        request_type_name='IamPoliciesCreatePolicyRequest',
        response_type_name='GoogleLongrunningOperation',
        supports_download=False,
    )

    def Delete(self, request, global_params=None):
      r"""Deletes a policy. This action is permanent.

      Args:
        request: (IamPoliciesDeleteRequest) input message
        global_params: (StandardQueryParameters, default: None) global arguments
      Returns:
        (GoogleLongrunningOperation) The response message.
      """
      config = self.GetMethodConfig('Delete')
      return self._RunMethod(
          config, request, global_params=global_params)

    Delete.method_config = lambda: base_api.ApiMethodInfo(
        flat_path='v2/policies/{policiesId}/{policiesId1}/{policiesId2}',
        http_method='DELETE',
        method_id='iam.policies.delete',
        ordered_params=['name'],
        path_params=['name'],
        query_params=['etag'],
        relative_path='v2/{+name}',
        request_field='',
        request_type_name='IamPoliciesDeleteRequest',
        response_type_name='GoogleLongrunningOperation',
        supports_download=False,
    )

    def Get(self, request, global_params=None):
      r"""Gets a policy.

      Args:
        request: (IamPoliciesGetRequest) input message
        global_params: (StandardQueryParameters, default: None) global arguments
      Returns:
        (GoogleIamV2Policy) The response message.
      """
      config = self.GetMethodConfig('Get')
      return self._RunMethod(
          config, request, global_params=global_params)

    Get.method_config = lambda: base_api.ApiMethodInfo(
        flat_path='v2/policies/{policiesId}/{policiesId1}/{policiesId2}',
        http_method='GET',
        method_id='iam.policies.get',
        ordered_params=['name'],
        path_params=['name'],
        query_params=[],
        relative_path='v2/{+name}',
        request_field='',
        request_type_name='IamPoliciesGetRequest',
        response_type_name='GoogleIamV2Policy',
        supports_download=False,
    )

    def ListPolicies(self, request, global_params=None):
      r"""Retrieves the policies of the specified kind that are attached to a resource. The response lists only policy metadata. In particular, policy rules are omitted.

      Args:
        request: (IamPoliciesListPoliciesRequest) input message
        global_params: (StandardQueryParameters, default: None) global arguments
      Returns:
        (GoogleIamV2ListPoliciesResponse) The response message.
      """
      config = self.GetMethodConfig('ListPolicies')
      return self._RunMethod(
          config, request, global_params=global_params)

    ListPolicies.method_config = lambda: base_api.ApiMethodInfo(
        flat_path='v2/policies/{policiesId}/{policiesId1}',
        http_method='GET',
        method_id='iam.policies.listPolicies',
        ordered_params=['parent'],
        path_params=['parent'],
        query_params=['pageSize', 'pageToken'],
        relative_path='v2/{+parent}',
        request_field='',
        request_type_name='IamPoliciesListPoliciesRequest',
        response_type_name='GoogleIamV2ListPoliciesResponse',
        supports_download=False,
    )

    def Update(self, request, global_params=None):
      r"""Updates the specified policy. You can update only the rules and the display name for the policy. To update a policy, you should use a read-modify-write loop: 1. Use GetPolicy to read the current version of the policy. 2. Modify the policy as needed. 3. Use `UpdatePolicy` to write the updated policy. This pattern helps prevent conflicts between concurrent updates.

      Args:
        request: (GoogleIamV2Policy) input message
        global_params: (StandardQueryParameters, default: None) global arguments
      Returns:
        (GoogleLongrunningOperation) The response message.
      """
      config = self.GetMethodConfig('Update')
      return self._RunMethod(
          config, request, global_params=global_params)

    Update.method_config = lambda: base_api.ApiMethodInfo(
        flat_path='v2/policies/{policiesId}/{policiesId1}/{policiesId2}',
        http_method='PUT',
        method_id='iam.policies.update',
        ordered_params=['name'],
        path_params=['name'],
        query_params=[],
        relative_path='v2/{+name}',
        request_field='<request>',
        request_type_name='GoogleIamV2Policy',
        response_type_name='GoogleLongrunningOperation',
        supports_download=False,
    )

  class V2Service(base_api.BaseApiService):
    """Service class for the v2 resource."""

    _NAME = 'v2'

    def __init__(self, client):
      super(IamV2.V2Service, self).__init__(client)
      self._upload_configs = {
          }

    def GetEffectivePolicies(self, request, global_params=None):
      r"""Deprecated: Use ListApplicablePolicies instead.

      Args:
        request: (IamGetEffectivePoliciesRequest) input message
        global_params: (StandardQueryParameters, default: None) global arguments
      Returns:
        (GoogleIamV2GetEffectivePoliciesResponse) The response message.
      """
      config = self.GetMethodConfig('GetEffectivePolicies')
      return self._RunMethod(
          config, request, global_params=global_params)

    GetEffectivePolicies.method_config = lambda: base_api.ApiMethodInfo(
        flat_path='v2/{v2Id}:getEffectivePolicies',
        http_method='GET',
        method_id='iam.getEffectivePolicies',
        ordered_params=['attachmentPoint'],
        path_params=['attachmentPoint'],
        query_params=['filter'],
        relative_path='v2/{+attachmentPoint}:getEffectivePolicies',
        request_field='',
        request_type_name='IamGetEffectivePoliciesRequest',
        response_type_name='GoogleIamV2GetEffectivePoliciesResponse',
        supports_download=False,
    )

    def ListApplicablePolicies(self, request, global_params=None):
      r"""Retrieves all the policies that are attached to the specified resource, or anywhere in the ancestry of the resource. For example, for a project this endpoint would return all the `denyPolicy` kind policies attached to the project, its parent folder (if any), and its parent organization (if any). The endpoint requires the same permissions that it would take to call `ListPolicies` or `GetPolicy`. The main reason to use this endpoint is as a policy admin to debug access issues for a resource.

      Args:
        request: (IamListApplicablePoliciesRequest) input message
        global_params: (StandardQueryParameters, default: None) global arguments
      Returns:
        (GoogleIamV2ListApplicablePoliciesResponse) The response message.
      """
      config = self.GetMethodConfig('ListApplicablePolicies')
      return self._RunMethod(
          config, request, global_params=global_params)

    ListApplicablePolicies.method_config = lambda: base_api.ApiMethodInfo(
        flat_path='v2/{v2Id}:listApplicablePolicies',
        http_method='GET',
        method_id='iam.listApplicablePolicies',
        ordered_params=['attachmentPoint'],
        path_params=['attachmentPoint'],
        query_params=['filter', 'pageSize', 'pageToken'],
        relative_path='v2/{+attachmentPoint}:listApplicablePolicies',
        request_field='',
        request_type_name='IamListApplicablePoliciesRequest',
        response_type_name='GoogleIamV2ListApplicablePoliciesResponse',
        supports_download=False,
    )
