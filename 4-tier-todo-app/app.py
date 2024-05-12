import streamlit as st
import boto3
import pandas as pd
import json
import datetime
import time


region = "us-west-2"
bedrock = boto3.client('bedrock-runtime',region_name=region)
textract = boto3.client('textract', region_name=region)

#Call Amazon Bedrock
def call_bedrock(bodyText, img_base64 = None):
    global region
    global bedrock
    #global bedrock2
    #Put your model name in the below string
    sonnet = "anthropic.claude-3-sonnet-20240229-v1:0"
    claude2 = "anthropic.claude-v2:1"

    modelId = sonnet
    
    #body = json.dumps({"inputText": bodyText,
    #                   "textGenerationConfig":{"maxTokenCount":2048, "temperature":0, "topP":0.9}})
    # "anthropic_version": "bedrock-2023-05-31"
    body = json.dumps({"prompt": bodyText, "max_tokens_to_sample": 500, "temperature": 0.1, "top_p": 0.1, "top_k": 10 })
    accept = 'application/json'
    contentType = 'application/json'
    contentBack = False
    backoff = 10
    retries = 0
    while contentBack != True:
        #Region Shifting if current region is "us-west-2"
        if (region == "us-west-2"):
            region = "us-east-1"
            bedrock = boto3.client('bedrock-runtime',region_name=region)
            #bedrock2 = boto3.client('bedrock',region_name=region)
        else:
            region = "us-west-2"
            bedrock = boto3.client('bedrock-runtime',region_name=region)
            #bedrock2 = boto3.client('bedrock',region_name=region)           
        print(f"region = {region}")
        print(f"retries={retries}")
        retries=retries+1
        if retries>3:
            print ("ERROR: timeout")
            return "Timeout: It was not possible to generate a response"
        try:
            response = ""
            if modelId == claude2:
                response = bedrock.invoke_model(body=body, modelId=modelId, accept=accept,contentType=contentType)
            else:
                if img_base64:
                    messages=[{ "role":'user', "content":[{"type": "image","source": {"type": "base64","media_type": "image/jpeg","data": img_base64}},{'type':'text','text': bodyText}]}]
                else:
                    messages=[{ "role":'user', "content":[{'type':'text','text': bodyText}]}]
                body = json.dumps({"max_tokens": 1000, "temperature": 0.1, "top_p": 0.1, "top_k": 10, "anthropic_version": "bedrock-2023-05-31", "messages":messages})
                response = bedrock.invoke_model(body=body, modelId=modelId)
            print(response)
            response_body = json.loads(response.get('body').read())
            ct = datetime.datetime.now()
            print(ct, " ", response_body)
  
            contentBack = True
            result_text = response_body['content'][0]['text']
            return result_text
            
        except Exception as error:    # Put your error handling logic here
            print("Model exception, will retry")
            print(bodyText)
            ct = datetime.datetime.now()
            print(ct, " ", error)
        time.sleep(backoff)
        backoff = backoff * 1.2

st.header("Data Visualization")

#shows a button to tupload a file
st.write(call_bedrock("What is the capital city of France?"))

uploaded_file= st.file_uploader("Upload a CSV file", type="csv")

#shows the data in the csv file in streamlit
if uploaded_file is not None:
    df= pd.read_csv(uploaded_file)
    st.write(df)    





