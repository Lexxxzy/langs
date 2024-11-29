from src.module import global_function, MyClass, function_with_local_variable

def class_obj_thing():
    obj = MyClass("I am an instance variable")
    obj.display_instance_variable()
    MyClass.display_class_variable()

    obj.static_method_example()


def main():
    global_function()

    class_obj_thing()

    MyClass.static_method_example()

    function_with_local_variable()
