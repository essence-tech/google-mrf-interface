"""Provide client for the Google MRF service."""

import grpc

from google_mrf_pb2 import GoogleMRFStub, Empty, Query


def get_client(address, pem_file=None):
    """Get googlemrf client."""
    if pem_file is not None:
        cred_file = open(pem_file).read()
        creds = grpc.ssl_channel_credentials(cred_file)
        channel = grpc.secure_channel(address, creds)
        return GoogleMRFStub(channel)
    else:
        channel = grpc.insecure_channel(address)
        return GoogleMRFStub(channel)
