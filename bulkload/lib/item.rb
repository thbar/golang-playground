class Item < ActiveRecord::Base
  attr_taggable :tags
  
  validates_presence_of :guid, :some_field, :some_other_field
end