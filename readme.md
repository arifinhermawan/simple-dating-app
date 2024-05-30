# Code Structure

The structure of this project is a mix of clean code and DDD. The structure itself is pretty straighforward. Basically the aim of this style is to emphasize SOLID and separation of layers.

The layer hierarchy goes like this:

Handler -> usecase -> service -> repo

This structure is pretty decent. It's easy to understand the purpose of each layer at a glance. But I know it's not perfect. So feel free to reach me out and give me feedback or suggestion 😄

## Folders

![alt text](image.png)

- cmd
- files
- internal

### cmd

cmd's sole purpose is to hold file to start the app.

### files

files is where we store config and other additional file that relates with development.

### internal

internal is the folder where all the magic happens.

## Internal Folders

### app

app is basically like a glue. It initialize the whole layers and it also stores methods that can be used by the whole layer in infrastructure.

So what's the point of infra anyway? why bother with having methods that can be used anywhere? Well, the methods that is stored in infrastructure are methods that commonly used in the code such as `json.Marshal`, `json.Unmarshal`, `time.Now`, and even metrics.

Storing those common methods in infra helps a lot with unit testing. Why? because by doing so we make unit testing much much easier because we can mock them. Sure we can do monkey patch as well. But we all know monkey patch is not very safe.

### handler

Handler is the layer that "talks" to the client. It's main job is:

1. parse input into usecase parameter
2. validate the input before sending them to usecase
3. parse output from usecase into presentation data structure

if done correctly, handler should not have any logic other that validation.

### usecase

Usecase is where the main logic lies. It contains business rule/features and it has no specific implementation of presentation or domain. Think about how you talk to Product team when you're discussing a project. Product team should be able to understand the logic in the usecase layer.

### service

Service is the layer where the heavy lifting happens. The implementation is getting more specific in this layer (i.e: fetching data from db then storing it in cache). It interacts with repo and does all the logic needed to achieve what's needed in usecase.

### repository

This is the layer where specific driver implementation lies. In this layer we can interact with data layer and also external API.

# How to run locally

## pre-requisites

1. go 1.17 or higher
2. postgresql

## steps to run

1. `make dep` (fetching all dependency)
2. `make run`
